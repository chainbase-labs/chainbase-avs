package config

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

// Config contains all of the configuration information for a coordinators.
type Config struct {
	EcdsaPrivateKey *ecdsa.PrivateKey `json:"-"`
	Logger          sdklogging.Logger `json:"-"`
	// we need the url for the eigensdk currently... eventually standardize api so as to
	// only take an ethClient or an rpcUrl (and build the ethClient at each constructor site)
	EthHttpRpcUrl               string
	EthWsRpcUrl                 string
	EthHttpClient               eth.Client `json:"-"`
	EthWsClient                 eth.Client `json:"-"`
	OperatorStateRetrieverAddr  common.Address
	RegistryCoordinatorAddr     common.Address
	CoordinatorServerIpPortAddr string
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement marshalJson
	SignerFn            signerv2.SignerFn `json:"-"`
	TxMgr               txmgr.TxManager   `json:"-"`
	CoordinatorAddress  common.Address
	FlinkGatewayHttpUrl string
	OssAccessKeyId      string `json:"-"`
	OssAccessKeySecret  string `json:"-"`
	TaskChains          []string
}

// ConfigRaw These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                 sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                   string              `yaml:"eth_rpc_url"`
	EthWsUrl                    string              `yaml:"eth_ws_url"`
	EcdsaPrivateKeyStorePath    string              `yaml:"ecdsa_private_key_store_path"`
	RegistryCoordinatorAddr     string              `yaml:"registry_coordinator_addr"`
	OperatorStateRetrieverAddr  string              `yaml:"operator_state_retriever_addr"`
	CoordinatorServerIpPortAddr string              `yaml:"coordinator_server_ip_port_address"`
	FlinkGatewayHttpUrl         string              `yaml:"flink_gateway_http_url"`
	OssAccessKeyId              string              `yaml:"oss_access_key_id"`
	OssAccessKeySecret          string              `yaml:"oss_access_key_secret"`
	TaskChains                  []string            `yaml:"task_chains"`
}

// NewConfig parses config file to read from from flags or environment variables
// Note: This config is shared by challenger and coordinator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context) (*Config, error) {
	var configRaw ConfigRaw
	configFilePath := ctx.GlobalString(ConfigFileFlag.Name)
	if configFilePath != "" {
		err := sdkutils.ReadYamlConfig(configFilePath, &configRaw)
		if err != nil {
			return nil, err
		}
	}

	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		return nil, err
	}

	ethRpcClient, err := eth.NewClient(configRaw.EthRpcUrl)
	if err != nil {
		logger.Error("Cannot create http ethClient", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(configRaw.EthWsUrl)
	if err != nil {
		logger.Error("Cannot create ws ethClient", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("COORDINATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Printf("COORDINATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}
	ecdsaPrivateKey, err := sdkecdsa.ReadKey(
		configRaw.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		logger.Error("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}

	coordinatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		logger.Error("Cannot get operator address", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, coordinatorAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, coordinatorAddr)

	config := &Config{
		EcdsaPrivateKey:             ecdsaPrivateKey,
		Logger:                      logger,
		EthWsRpcUrl:                 configRaw.EthWsUrl,
		EthHttpRpcUrl:               configRaw.EthRpcUrl,
		EthHttpClient:               ethRpcClient,
		EthWsClient:                 ethWsClient,
		OperatorStateRetrieverAddr:  common.HexToAddress(configRaw.OperatorStateRetrieverAddr),
		RegistryCoordinatorAddr:     common.HexToAddress(configRaw.RegistryCoordinatorAddr),
		CoordinatorServerIpPortAddr: configRaw.CoordinatorServerIpPortAddr,
		SignerFn:                    signerV2,
		TxMgr:                       txMgr,
		CoordinatorAddress:          coordinatorAddr,
		FlinkGatewayHttpUrl:         configRaw.FlinkGatewayHttpUrl,
		OssAccessKeyId:              configRaw.OssAccessKeyId,
		OssAccessKeySecret:          configRaw.OssAccessKeySecret,
		TaskChains:                  configRaw.TaskChains,
	}
	config.validate()
	return config, nil
}

func (c *Config) validate() {
	if c.EcdsaPrivateKey == nil {
		panic("Config: EcdsaPrivateKey is required")
	}
	if c.Logger == nil {
		panic("Config: Logger is required")
	}
	if c.EthHttpRpcUrl == "" {
		panic("Config: EthHttpRpcUrl is required")
	}
	if c.EthWsRpcUrl == "" {
		panic("Config: EthWsRpcUrl is required")
	}
	if c.EthHttpClient == nil {
		panic("Config: EthHttpClient is required")
	}
	if c.EthWsClient == nil {
		panic("Config: EthWsClient is required")
	}
	if c.OperatorStateRetrieverAddr == common.HexToAddress("") {
		panic("Config: OperatorStateRetrieverAddr is required")
	}
	if c.RegistryCoordinatorAddr == common.HexToAddress("") {
		panic("Config: RegistryCoordinatorAddr is required")
	}
	if c.CoordinatorServerIpPortAddr == "" {
		panic("Config: CoordinatorServerIpPortAddr is required")
	}
	if c.SignerFn == nil {
		panic("Config: SignerFn is required")
	}
	if c.TxMgr == nil {
		panic("Config: TxMgr is required")
	}
	if c.CoordinatorAddress == common.HexToAddress("") {
		panic("Config: CoordinatorAddress is required")
	}
}

var (
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	/* Optional Flags */
)

var requiredFlags = []cli.Flag{
	ConfigFileFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
