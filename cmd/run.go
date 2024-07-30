package cmd

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

	"github.com/chainbase-avs/cli/pkg/prometheus"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "monitor manualscript requests and send signed to gateway",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		avsAddr, err := Run(cmd.Context())
		if err != nil {
			slog.Error("failed to run", "error", err)
			os.Exit(1)
		}

		router := httprouter.New()
		router.GET("/eigen/node/health", handleHealth)

		router.Handler("GET", "/metrics", promhttp.Handler())

		go func() {

			for {

				prometheus.UpdateHostMetrics(avsAddr)
				slog.Info("update host metrics", "avsAddr", avsAddr, "ip", prometheus.GetOutboundIP())
				time.Sleep(15 * time.Second)
			}

		}()

		err = http.ListenAndServe(fmt.Sprintf(":%d", HealthCheckPort), router)
		if err != nil {
			log.Println(err)
		}
	},
}

func handleHealth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func Run(ctx context.Context) (string, error) {

	var operatorAddress string

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
		return operatorAddress, err
	} else if err := deps.VerifyFunc(eigenCfg.Operator); err != nil {
		return operatorAddress, err
	}

	password := viper.GetString(KeystorePassword)

	privateKey, err := eigenecdsa.ReadKey(eigenCfg.PrivateKeyStorePath, password)
	if err != nil {
		return operatorAddress, err
	}

	//1. eth client
	client, err := ethclient.Dial(viper.GetString(RPC_URL))
	if err != nil {
		slog.Error("failed to connect to the Ethereum client", "error", err)
		return operatorAddress, err
	}

	//2. create contract binding avsInstance
	avsInstance, err := mc.NewIAVS(contractAddress, client)
	if err != nil {
		slog.Error("failed to create a AVS binding instance", "error", err)
		return operatorAddress, err
	}

	//3. call avs check if registered
	operatorList, err := avsInstance.Operators(&bind.CallOpts{})
	if err != nil {
		slog.Error("failed to check if AVS is registered", "error", err)
		return operatorAddress, err
	}

	for _, operator := range operatorList {
		if operator == crypto.PubkeyToAddress(privateKey.PublicKey) {
			slog.Info("operator", "address", operator.Hex())
			slog.Info("AVS is registered,continue")
			return operator.Hex(), nil
		}

	}
	return AVSContractAddress, errors.New("AVS is not registered")
}
