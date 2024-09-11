package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	mc "github.com/chainbase-labs/chainbase-avs/contracts/bindings/deprecated"
	"github.com/chainbase-labs/chainbase-avs/pkg/prometheus"
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

		appRouter := httprouter.New()
		appRouter.GET("/eigen/node/health", handleHealth)

		metricsRouter := httprouter.New()
		metricsRouter.Handler("GET", "/metrics", promhttp.Handler())

		go func() {

			for {

				prometheus.UpdateHostMetrics(avsAddr)
				slog.Info("update host metrics", "avsAddr", avsAddr, "ip", prometheus.GetOutboundIP(), "job_manager_status", prometheus.GetFlinkJobManagerStatus())
				time.Sleep(15 * time.Second)
			}

		}()

		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt(NodeAppPort)), appRouter)
			if err != nil {
				slog.Error("Failed to start app server", "error", err)
			}
		}()

		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt(NodeMetricsPort)), metricsRouter)
			if err != nil {
				slog.Error("Failed to start metrics server", "error", err)
			}
		}()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigChan
		slog.Info("Received signal, shutting down", "signal", sig)

	},
}

func handleHealth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func Run(ctx context.Context) (string, error) {

	var operatorAddress string

	contractAddress := common.HexToAddress(viper.GetString(AVSContractAddress))

	//0.decode keystore get ecdsa private key
	privateKey, err := eigenecdsa.ReadKey(viper.GetString(OperatorKeystorePath), viper.GetString(KeystorePassword))
	if err != nil {
		return operatorAddress, err
	}

	//1. eth client
	client, err := ethclient.Dial(viper.GetString(NodeChainRpc))
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
