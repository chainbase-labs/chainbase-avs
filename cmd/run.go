package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/pkg/errors"

	mc "github.com/chainbase-avs/cli/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "monitor manualscript requests and send signed to gateway",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := Run(cmd.Context())
		if err != nil {
			slog.Error("failed to run", "error", err)
			os.Exit(1)
		}
	},
}

func Run(ctx context.Context) error {

	deps := RegDeps{
		Prompter: eigenutils.NewPrompter(),
		VerifyFunc: func(op eigensdktypes.Operator) error {
			return op.Validate()
		},
	}

	contractAddress := common.HexToAddress(viper.GetString(AVSContractAddress))

	//0.read eigenlayer config to get ecdsa private key
	eigenCfg, err := readConfig(viper.GetString(OperatorConfigPath))
	if err != nil {
		return err
	} else if err := deps.VerifyFunc(eigenCfg.Operator); err != nil {
		return err
	}

	password := viper.GetString(KeystorePassword)

	privateKey, err := eigenecdsa.ReadKey(eigenCfg.PrivateKeyStorePath, password)
	if err != nil {
		return err
	}

	//1. eth client
	client, err := ethclient.Dial(viper.GetString(RPC_URL))
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

	//3. call avs check if registered
	operatorList, err := avsInstance.Operators(&bind.CallOpts{})
	if err != nil {
		slog.Error("failed to check if AVS is registered", "error", err)
		return err
	}

	for _, operator := range operatorList {
		slog.Info("operator", "address", operator.Hex())
		if operator == crypto.PubkeyToAddress(privateKey.PublicKey) {
			slog.Info("AVS is registered,continue")
			return nil
		}

	}
	return errors.New("AVS is not registered")
}
