package node

import (
	"context"
	"fmt"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/core"
	"github.com/chainbase-labs/chainbase-avs/core/chainio"
	"github.com/chainbase-labs/chainbase-avs/metrics"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

const AvsName = "chainbase"
const SemVer = "0.0.1"

type ManuscriptNode struct {
	nodepb.UnimplementedManuscriptNodeServiceServer
	config           types.NodeConfig
	logger           logging.Logger
	ethClient        eth.Client
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.IAvsReader
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *bindings.ChainbaseServiceManagerNewTaskCreated
	// ip address of coordinator
	coordinatorServerIpPortAddr string
	// nodeServerIpPortAddr is the IP address and port of the Node gRPC server.
	// This public address is submitted to the contract during the RegisterOperatorWithAvs and can be requested by the coordinator.
	nodeServerIpPortAddr string
	// rpc client to send signed task responses to coordinator
	coordinatorRpcClient CoordinatorRpcClienter
}

func NewNodeFromConfig(c types.NodeConfig) (*ManuscriptNode, error) {
	var logLevel logging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AvsName, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AvsName, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AvsName, SemVer, c.NodeApiIpPortAddress, logger)

	var ethRpcClient eth.Client
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AvsName, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Error("Cannot create http eth client", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Error("Cannot create http eth client", "err", err)
			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Error("Cannot parse bls private key", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}
	chainConfig := clients.BuildAllConfig{
		EthHttpUrl:                 c.EthRpcUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
		AvsName:                    AvsName,
		PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))

	avsWriter, err := chainio.BuildAvsWriter(
		txMgr, common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AvsName, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	coordinatorRpcClient, err := NewCoordinatorRpcClient(c.CoordinatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create CoordinatorRpcClient. Is coordinator running?", "err", err)
		return nil, err
	}

	msNode := &ManuscriptNode{
		config:                      c,
		logger:                      logger,
		metricsReg:                  reg,
		metrics:                     avsAndEigenMetrics,
		nodeApi:                     nodeApi,
		ethClient:                   ethRpcClient,
		avsWriter:                   avsWriter,
		avsReader:                   avsReader,
		eigenlayerReader:            sdkClients.ElChainReader,
		eigenlayerWriter:            sdkClients.ElChainWriter,
		blsKeypair:                  blsKeyPair,
		operatorAddr:                common.HexToAddress(c.OperatorAddress),
		coordinatorServerIpPortAddr: c.CoordinatorServerIpPortAddress,
		nodeServerIpPortAddr:        c.NodeServerIpPortAddress,
		coordinatorRpcClient:        coordinatorRpcClient,
		newTaskCreatedChan:          make(chan *bindings.ChainbaseServiceManagerNewTaskCreated),
		operatorId:                  [32]byte{0}, // this is set below
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, msNode.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	msNode.operatorId = operatorId
	logger.Info("ManuscriptNode info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", msNode.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", msNode.blsKeypair.GetPubKeyG2(),
	)

	return msNode, nil

}

func (n *ManuscriptNode) Start(ctx context.Context) error {
	operatorIsRegistered, err := n.avsReader.IsOperatorRegistered(&bind.CallOpts{}, n.operatorAddr)
	if err != nil {
		n.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	n.logger.Infof("Starting operator.")

	go func() {
		err = n.startServer(ctx)
		if err != nil {
			n.logger.Error("Starting manuscript node rpc server error", "err", err)
		}
	}()

	if n.config.EnableNodeApi {
		n.nodeApi.Start()
	}
	var metricsErrChan <-chan error
	if n.config.EnableMetrics {
		metricsErrChan = n.metrics.Start(ctx, n.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			n.logger.Fatal("Error in metrics server", "err", err)
		case newTaskCreatedLog := <-n.newTaskCreatedChan:
			n.metrics.IncNumTasksReceived()
			taskResponse := n.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			signedTaskResponse, err := n.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go n.coordinatorRpcClient.SendSignedTaskResponseToCoordinator(signedTaskResponse)
		}
	}
}

// ProcessNewTaskCreatedLog Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (n *ManuscriptNode) ProcessNewTaskCreatedLog(newTaskCreatedLog *bindings.ChainbaseServiceManagerNewTaskCreated) *bindings.IChainbaseServiceManagerTaskResponse {
	n.logger.Debug("Received new task", "task", newTaskCreatedLog)
	n.logger.Info("Received new task",
		"taskDetails", newTaskCreatedLog.Task.TaskDetails,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)
	//TODO execute task
	response := "0x9c2643e05c22861a55cb4a5455678b5acec4e0c5d1c0311d22e2abacfa2bab29"
	taskResponse := &bindings.IChainbaseServiceManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		TaskResponse:       response,
	}
	return taskResponse
}

func (n *ManuscriptNode) SignTaskResponse(taskResponse *bindings.IChainbaseServiceManagerTaskResponse) (*coordinatorpb.SignedTaskResponseRequest, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		n.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}

	blsSignature := n.blsKeypair.SignMessage(taskResponseHash)

	signedTaskResponse := &coordinatorpb.SignedTaskResponseRequest{
		TaskResponse: &coordinatorpb.IChainbaseServiceManagerTaskResponse{
			ReferenceTaskIndex: taskResponse.ReferenceTaskIndex,
			TaskResponse:       taskResponse.TaskResponse,
		},
		BlsSignature: &coordinatorpb.Signature{
			G1Point: &coordinatorpb.G1Point{
				X: blsSignature.X[:],
				Y: blsSignature.Y[:],
			},
		},
		OperatorId: n.operatorId[:],
	}

	n.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
