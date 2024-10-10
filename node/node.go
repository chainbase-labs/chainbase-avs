package node

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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
	"github.com/cbergoon/merkletree"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	dockerClient "github.com/docker/docker/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"

	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/core"
	"github.com/chainbase-labs/chainbase-avs/core/chainio"
	"github.com/chainbase-labs/chainbase-avs/node/metrics"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

const AvsName = "chainbase_avs"
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
	// receive task response from ProcessNewTaskCreatedLog
	TaskResponseChan chan *bindings.IChainbaseServiceManagerTaskResponse
	// ip address of coordinator
	coordinatorServerIpPortAddr string
	// node grpc server address
	nodeGrpcServerAddress string
	// nodeServerIpPortAddr is the IP address and port of the Node gRPC server.
	// This public address is submitted to the contract during the RegisterOperatorWithAvs and can be requested by the coordinator.
	nodeSocket string
	// rpc client to send signed task responses to coordinator
	coordinatorRpcClient CoordinatorRpcClienter
	// task flink jobID
	taskJobIDs map[types.TaskIndex]string
	// docker client
	dockerClient *dockerClient.Client
	// postgres db
	db *sql.DB
	// job manager host
	jobManagerHost string
	// job manager port
	jobManagerPort string
}

func NewNodeFromConfig(c types.NodeConfig, cliCommand bool) (*ManuscriptNode, error) {
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

	cli, err := dockerClient.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Error("Cannot create docker client", "err", err)
		return nil, err
	}

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost, c.PostgresPort, c.PostgresUser, c.PostgresPassword, c.PostgresDatabase)

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
		nodeGrpcServerAddress:       c.NodeGrpcServerAddress,
		nodeSocket:                  c.NodeSocket,
		coordinatorRpcClient:        coordinatorRpcClient,
		newTaskCreatedChan:          make(chan *bindings.ChainbaseServiceManagerNewTaskCreated),
		TaskResponseChan:            make(chan *bindings.IChainbaseServiceManagerTaskResponse),
		operatorId:                  [32]byte{0}, // this is set below
		taskJobIDs:                  make(map[types.TaskIndex]string),
		dockerClient:                cli,
		jobManagerHost:              c.JobManagerHost,
		jobManagerPort:              c.JobManagerPort,
	}

	if !cliCommand {
		db, err := sql.Open("postgres", dataSource)
		if err != nil {
			log.Fatal("Error connecting to the database: ", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Error pinging the database: ", err)
		}

		msNode.db = db
		fmt.Println("Successfully connected to the database")
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, msNode.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	msNode.operatorId = operatorId
	logger.Info("ManuscriptNode info",
		"operatorId", hex.EncodeToString(operatorId[:]),
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
		return fmt.Errorf("operator is not registered. Registering operator using the chainbase-cli before starting operator")
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
		go func() {
			for {
				ip, jobManagerStatus := n.metrics.UpdateNodeMetrics(n.operatorAddr.String(), AvsName, n.jobManagerHost, n.jobManagerPort)
				n.logger.Info("update host metrics", "operator address", n.operatorAddr.String(), "ip", ip, "job_manager_status", jobManagerStatus)
				time.Sleep(1 * time.Minute)
			}

		}()
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
			n.metrics.IncNumTaskReceived()
			go n.ProcessNewTaskCreatedLog(newTaskCreatedLog)
		case taskResponse := <-n.TaskResponseChan:
			signedTaskResponse, err := n.SignTaskResponse(taskResponse)
			if err != nil {
				n.metrics.IncNumTaskFailed()
				continue
			}
			go n.coordinatorRpcClient.SendSignedTaskResponseToCoordinator(signedTaskResponse)
		}
	}
}

// ProcessNewTaskCreatedLog Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct to TaskResponseChan channel.
func (n *ManuscriptNode) ProcessNewTaskCreatedLog(newTaskCreatedLog *bindings.ChainbaseServiceManagerNewTaskCreated) {
	n.logger.Info("Received new task",
		"taskDetails", newTaskCreatedLog.Task.TaskDetails,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	parsedTaskDetails, err := core.ParseTaskDetails(newTaskCreatedLog.Task.TaskDetails)
	if err != nil {
		n.metrics.IncNumTaskFailed()
		n.logger.Error("Failed to parse task details", "err", err)
		return
	}

	executeStartTime := time.Now()

	if err := n.ExecuteTask(newTaskCreatedLog.TaskIndex, parsedTaskDetails); err != nil {
		n.metrics.IncNumTaskFailed()
		n.logger.Error("Failed to execute task", "err", err)
		return
	}

	if err := n.WaitTaskCompletion(newTaskCreatedLog.TaskIndex, parsedTaskDetails); err != nil {
		n.metrics.IncNumTaskFailed()
		n.logger.Error("Error wait task completion", "err", err)
		n.CancelTaskJob(newTaskCreatedLog.TaskIndex)
		return
	}

	executeEndTime := time.Now()
	executeDuration := executeEndTime.Sub(executeStartTime)
	n.metrics.ObserveTaskExecutionTime(executeDuration.Minutes())

	response, err := n.QueryTaskResponse(newTaskCreatedLog.TaskIndex)
	if err != nil {
		n.metrics.IncNumTaskFailed()
		n.logger.Error("Error query task response", "err", err)
		n.CancelTaskJob(newTaskCreatedLog.TaskIndex)
		return
	}

	n.CancelTaskJob(newTaskCreatedLog.TaskIndex)

	n.TaskResponseChan <- &bindings.IChainbaseServiceManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		TaskResponse:       response,
	}
}

func (n *ManuscriptNode) ExecuteTask(taskIndex uint32, taskDetails *core.TaskDetails) error {
	n.logger.Info("Executing task", "task index", taskIndex, "task details", taskDetails)

	containerName := "chainbase_jobmanager"
	cmd := []string{
		"/bin/sh",
		"-c",
		fmt.Sprintf("./bin/flink run /opt/proof/proof-1.0-SNAPSHOT.jar %s %d %d %d %d",
			taskDetails.Chain, taskDetails.StartBlock, taskDetails.EndBlock, taskDetails.Difficulty, taskIndex),
	}
	execConfig := dockerTypes.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	}

	execID, err := n.dockerClient.ContainerExecCreate(context.Background(), containerName, execConfig)
	if err != nil {
		n.logger.Error("Failed to create exec instance", "error", err)
		return err
	}

	resp, err := n.dockerClient.ContainerExecAttach(context.Background(), execID.ID, dockerTypes.ExecStartCheck{})
	if err != nil {
		n.logger.Error("Failed to attach to exec instance", "error", err)
		return err
	}
	defer resp.Close()

	var output bytes.Buffer
	_, err = io.Copy(&output, resp.Reader)
	if err != nil {
		n.logger.Error("Failed to read exec output", "error", err)
		return err
	}

	jobID, err := extractJobID(output.String())
	if err != nil {
		n.logger.Error("Error extract job id", "err", err)
		return err
	}

	n.taskJobIDs[taskIndex] = jobID
	n.logger.Info("Job has been submitted", "jobID", jobID)
	return nil
}

func (n *ManuscriptNode) WaitTaskCompletion(taskIndex uint32, taskDetails *core.TaskDetails) error {
	retryInterval := time.Second * 30
	resultCount := taskDetails.EndBlock - taskDetails.StartBlock + 1
	for {
		if time.Now().Unix() >= taskDetails.Deadline {
			return errors.New("task was not completed on time")
		}
		query := `SELECT pow_result FROM pow_results WHERE task_index = $1`
		rows, err := n.db.Query(query, taskIndex)
		if err != nil {
			n.logger.Error("Error query task response", "err", err)
			continue
		}
		rowsCount := 0
		for rows.Next() {
			rowsCount++
		}
		rows.Close()
		// there is all required task result in postgres db
		if rowsCount >= resultCount {
			n.logger.Info("Task is completed", "taskIndex", taskIndex)
			return nil
		}
		// TODO print log about task execute progress
		time.Sleep(retryInterval)
	}
}

func (n *ManuscriptNode) QueryTaskResponse(taskIndex uint32) (string, error) {
	query := `SELECT pow_result FROM pow_results WHERE task_index = $1`
	rows, err := n.db.Query(query, taskIndex)
	if err != nil {
		n.logger.Error("Error query task response", "err", err)
		return "", err
	}
	defer rows.Close()

	resultContents := make([]merkletree.Content, 0)
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			n.logger.Error("Error scanning row", "err", err)
			return "", err
		}
		resultContents = append(resultContents, ResultContent{
			result: result,
		})
	}

	if err = rows.Err(); err != nil {
		n.logger.Error("Error iterating rows", "err", err)
		return "", err
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(resultContents)
	if err != nil {
		n.logger.Error("Error create merkle tree", "err", err)
		return "", err
	}

	//return the Merkle Root of the tree
	response := "0x" + hex.EncodeToString(t.MerkleRoot())
	return response, nil
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

func (n *ManuscriptNode) CancelTaskJob(taskIndex uint32) {
	jobID, ok := n.taskJobIDs[taskIndex]
	if !ok {
		n.logger.Error("Task JobID is not exist", "taskIndex", taskIndex)
	}
	n.logger.Info("Task Job  is cancelling", "taskIndex", taskIndex, "JobID", jobID)

	containerName := "chainbase_jobmanager"
	cmd := []string{"flink", "cancel", jobID}
	execConfig := dockerTypes.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	}

	execID, err := n.dockerClient.ContainerExecCreate(context.Background(), containerName, execConfig)
	if err != nil {
		n.logger.Error("Failed to create exec instance", "error", err)
		return
	}

	resp, err := n.dockerClient.ContainerExecAttach(context.Background(), execID.ID, dockerTypes.ExecStartCheck{})
	if err != nil {
		n.logger.Error("Failed to attach to exec instance", "error", err)
		return
	}
	defer resp.Close()

	n.logger.Info("Task job is cancelled", "taskIndex", taskIndex, "JobID", jobID)
}
