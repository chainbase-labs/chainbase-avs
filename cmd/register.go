package cmd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"cosmossdk.io/errors"
	mc "github.com/chainbase-avs/cli/bindings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	eigentypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

func init() {
	rootCmd.AddCommand(registerCmd)
	BindRegisterConfig(registerCmd, &cfg)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register to avs",
	Long:  `register to avs`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Register(cmd.Context(), cfg)
		if err != nil {
			slog.Error("failed to register", "error", err)
		}
	},
}

var (
	AVSContractAddress = "0x7F521D31170266E49B269B06C65555badEB2a195"
	AVSDirContractAddr = "0x055733000064333CaDDbC92763c58BF0192fFeBf"
	RPC_URL            = "https://rpc.ankr.com/eth_holesky"
)

// makeAuth creates a transaction signer from a private key
func makeAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	slog.Info("public key ", "address:", crypto.PubkeyToAddress(*publicKeyECDSA))

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}

// RegDeps contains the Register dependencies that are abstracted for testing.
type RegDeps struct {
	Prompter   eigenutils.Prompter
	VerifyFunc func(eigensdktypes.Operator) error
}

func Register(ctx context.Context, cfg RegConfig) error {

	deps := RegDeps{
		Prompter: eigenutils.NewPrompter(),
		VerifyFunc: func(op eigensdktypes.Operator) error {
			return op.Validate()
		},
	}

	contractAddress := common.HexToAddress(AVSContractAddress)
	avsDirAddr := common.HexToAddress(AVSDirContractAddr)

	//0.read eigenlayer config to get ecdsa private key
	eigenCfg, err := readConfig(cfg.ConfigFile)
	if err != nil {
		return err
	} else if err := deps.VerifyFunc(eigenCfg.Operator); err != nil {
		return err
	}

	password, err := deps.Prompter.InputHiddenString("Enter password to decrypt the ecdsa private key:", "",
		func(string) error {
			return nil
		},
	)
	if err != nil {
		return err
	}
	privateKey, err := eigenecdsa.ReadKey(eigenCfg.PrivateKeyStorePath, password)
	if err != nil {
		return err
	}

	//1. eth client
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		slog.Error("failed to connect to the Ethereum client", "error", err)
		return err
	}

	//2. create contract binding avsInstance
	avsInstance, err := mc.NewIAVS(contractAddress, client)
	if err != nil {
		slog.Error("failed to create a AVS binding instance", "error", err)
		return err
	}
	avsDirInstance, err := mc.NewIAVSDirectory(avsDirAddr, client)
	if err != nil {
		slog.Error("failed to create a AVSDir binding instance", "error", err)
		return err
	}

	//3. sign related
	auth, err := makeAuth(client, privateKey)
	if err != nil {
		slog.Error("failed to create a transaction signer", "error", err)
	}

	// 4. make vaild signature
	bytes := make([]byte, 32)
	rand.Read(bytes)
	salt := crypto.Keccak256Hash(bytes)                    // Salt can be anything, it should just be unique.
	expiry := big.NewInt(time.Now().Add(time.Hour).Unix()) // Sig is 1 Hour valid
	digestHash, err := avsDirInstance.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{},
		auth.From,       // operator address
		contractAddress, // AVS address
		salt,
		expiry)
	if err != nil {
		slog.Error("failed to calculate digest hash", "error", err)
	}
	digestHashSlice := digestHash[:]

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, digestHashSlice)
	if err != nil {
		log.Fatal(err)

	}
	v := new(big.Int).SetInt64(27) // v = 27 or 28 default:27
	signature := append(r.Bytes(), s.Bytes()...)
	signature = append(signature, byte(v.Uint64()))

	slog.Info("sign", "signature", signature)
	slog.Info("salt", "salt", salt)
	slog.Info("expiry", "expiry", expiry)

	// 5. register
	tx, err := avsInstance.RegisterOperator(auth, mc.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    expiry,
	})
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("tx hash", "hash", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		reason, err := getTxFailureReason(client, tx, receipt.BlockNumber)
		if err != nil {
			slog.Error("failed to get failure reason", "error", err)
		} else {
			slog.Error("failed to register operator", "reason", reason)
		}
	}

	return nil

}

func getTxFailureReason(client *ethclient.Client, tx *types.Transaction, blockNumber *big.Int) (string, error) {

	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return "", fmt.Errorf("failed to get tx sender: %v", err)
	}
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	// use CallContract mock tx
	result, err := client.CallContract(context.Background(), msg, blockNumber)
	if err != nil {
		return "", fmt.Errorf("call contract failed: %v", err)
	}

	return string(result), nil
}

// readConfig returns the eigen-layer operator configuration from the given file.
func readConfig(file string) (eigentypes.OperatorConfigNew, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "eigen config file not found")
	}

	bz, err := os.ReadFile(file)
	if err != nil {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "read eigen config file")
	}

	var config eigentypes.OperatorConfigNew
	if err := yaml.Unmarshal(bz, &config); err != nil {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "unmarshal eigen config file")
	}

	return config, nil
}
