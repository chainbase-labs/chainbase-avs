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

	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/chainbase-labs/chainbase-avs/contracts/bindings/deprecated"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register to avs",
	Long:  `register to avs`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Register(cmd.Context())
		if err != nil {
			slog.Error("failed to register", "error", err)
		}
	},
}

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

func Register(ctx context.Context) error {

	contractAddress := common.HexToAddress(viper.GetString(AVSContractAddress))
	avsDirAddr := common.HexToAddress(viper.GetString(AVSDirContractAddr))

	//0.read eigenlayer config to get ecdsa private key
	//TODO: what if is smart contract
	privateKey, err := eigenecdsa.ReadKey(viper.GetString(OperatorKeystorePath), viper.GetString(KeystorePassword))
	if err != nil {
		return err
	}

	//1. eth client
	client, err := ethclient.Dial(viper.GetString(NodeChainRpc))
	if err != nil {
		slog.Error("failed to connect to the Ethereum client", "error", err)
		return err
	}

	//2. create contract binding avsInstance
	avsInstance, err := deprecated.NewIAVS(contractAddress, client)
	if err != nil {
		slog.Error("failed to create a AVS binding instance", "error", err)
		return err
	}
	avsDirInstance, err := deprecated.NewIAVSDirectory(avsDirAddr, client)
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

	sig, err := crypto.Sign(digestHashSlice, privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign digest hash: %v", err)
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := sig[64] + 27 // add 27 to v to conform with the Ethereum standard

	signature := append(r.Bytes(), s.Bytes()...)
	signature = append(signature, v)

	slog.Debug("signature", "r", r, "s", s, "v", v)
	slog.Debug("salt", "salt", salt)
	slog.Debug("expiry", "expiry", expiry)

	// 5. register
	tx, err := avsInstance.RegisterOperator(auth, deprecated.ISignatureUtilsSignatureWithSaltAndExpiry{
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
