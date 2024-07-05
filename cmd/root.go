package cmd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	mc "github.com/chainbase-avs/cli/bindings" // 导入生成的合约包
)

var (
	cfg RegConfig
)

var rootCmd = &cobra.Command{
	Use:   "chianbase-avs",
	Short: "chianbase-avs",
	Long:  `chianbase-avs`,
	Run: func(cmd *cobra.Command, args []string) {
		Register(cmd.Context(), cfg)
	},
}

func init() {
	BindRegConfig(rootCmd, &cfg)
}

func Execute() error {
	return rootCmd.Execute()
}

// Register registers the operator with the ChainBase AVS contract.
//
// It assumes that the operator is already registered with the Eigen-Layer
// and that the eigen-layer configuration file (and ecdsa keystore) is present on disk.

var (
	AVSContractAddress = "0x7F521D31170266E49B269B06C65555badEB2a195"
	AVSDirContractAddr = "0x055733000064333CaDDbC92763c58BF0192fFeBf"
	RPC_URL            = "https://rpc.ankr.com/eth_holesky"
)

func Register(ctx context.Context, cfg RegConfig) error {

	contractAddress := common.HexToAddress(AVSContractAddress)
	avsDirAddr := common.HexToAddress(AVSDirContractAddr)
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		slog.Error("failed to convert private key", "error", err)
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
	auth, err := makeAuth(client, cfg)
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

	// 使用 CallContract 来模拟交易
	result, err := client.CallContract(context.Background(), msg, blockNumber)
	if err != nil {
		return "", fmt.Errorf("call contract failed: %v", err)
	}

	// 解析错误信息
	return string(result), nil
}

// makeAuth creates a transaction signer from a private key
func makeAuth(client *ethclient.Client, cfg RegConfig) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
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
	auth.Value = big.NewInt(0)     // 如果合约调用需要发送ETH，则这里设置
	auth.GasLimit = uint64(300000) // 设置一个合适的gas限制
	auth.GasPrice = gasPrice
	return auth, nil
}
