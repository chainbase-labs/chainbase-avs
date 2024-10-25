package coordinator

import (
	"context"
	"encoding/hex"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/prometheus/client_golang/prometheus"

	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/coordinator/metrics"
	"github.com/chainbase-labs/chainbase-avs/coordinator/types"
	"github.com/chainbase-labs/chainbase-avs/core"
	"github.com/chainbase-labs/chainbase-avs/core/chainio"
	"github.com/chainbase-labs/chainbase-avs/core/config"
)

const AvsName = "chainbase_coordinator"

// Coordinator sends tasks onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the coordinator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Coordinator struct {
	coordinatorpb.UnimplementedCoordinatorServiceServer
	logger           logging.Logger
	serverIpPortAddr string
	ethClient        eth.Client
	avsWriter        chainio.IAvsWriter
	avsSubscriber    chainio.IAvsSubscriber
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *bindings.ChainbaseServiceManagerNewTaskCreated
	// avs registry service
	avsRegistryService avsregistry.AvsRegistryService
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskIndex]bindings.IChainbaseServiceManagerTask
	tasksMu               sync.RWMutex
	taskResponses         map[types.TaskIndex]map[sdktypes.TaskResponseDigest]bindings.IChainbaseServiceManagerTaskResponse
	taskResponsesMu       sync.RWMutex
	flinkClient           *FlinkClient
	taskChains            []string
	taskDurationMinutes   int64
	metricsReg            *prometheus.Registry
	metrics               metrics.Metrics
	quorumThreshold       uint8
}

// NewCoordinator creates a new Coordinator with the provided config.
func NewCoordinator(c *config.Config) (*Coordinator, error) {
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsWriter", "err", err)
		return nil, err
	}

	avsSubscriber, err := chainio.BuildAvsSubscriber(c.RegistryCoordinatorAddr,
		c.OperatorStateRetrieverAddr, c.EthWsClient, c.Logger,
	)
	if err != nil {
		c.Logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	chainConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    AvsName,
		PromMetricsIpPortAddress:   c.CoordinatorMetricsIpPortAddress,
	}
	chainClients, err := clients.BuildAll(chainConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create sdk clients", "err", err)
		return nil, err
	}

	hashFunction := func(taskResponse sdktypes.TaskResponse) (sdktypes.TaskResponseDigest, error) {
		r := taskResponse.(bindings.IChainbaseServiceManagerTaskResponse)
		taskResponseDigest, err := core.GetTaskResponseDigest(&r)
		if err != nil {
			return sdktypes.TaskResponseDigest{}, err
		}

		return taskResponseDigest, nil
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), chainClients.AvsRegistryChainSubscriber, chainClients.AvsRegistryChainReader, nil, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, hashFunction, c.Logger)
	flinkClient := NewFlinkClient(c.FlinkGatewayHttpUrl, c.OssAccessKeyId, c.OssAccessKeySecret)

	if len(c.TaskChains) == 0 {
		return nil, errors.New("no available task chains")
	}

	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AvsName, c.CoordinatorMetricsIpPortAddress, reg, c.Logger)
	coordinatorMetrics := metrics.NewCoordinatorMetrics(AvsName, eigenMetrics, reg)

	return &Coordinator{
		logger:                c.Logger,
		serverIpPortAddr:      c.CoordinatorServerIpPortAddr,
		ethClient:             chainClients.EthHttpClient,
		avsWriter:             avsWriter,
		avsSubscriber:         avsSubscriber,
		newTaskCreatedChan:    make(chan *bindings.ChainbaseServiceManagerNewTaskCreated),
		avsRegistryService:    avsRegistryService,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskIndex]bindings.IChainbaseServiceManagerTask),
		taskResponses:         make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]bindings.IChainbaseServiceManagerTaskResponse),
		flinkClient:           flinkClient,
		taskChains:            c.TaskChains,
		taskDurationMinutes:   c.TaskDurationMinutes,
		metricsReg:            reg,
		metrics:               coordinatorMetrics,
		quorumThreshold:       c.QuorumThreshold,
	}, nil
}

func (c *Coordinator) Start(ctx context.Context) error {
	c.logger.Infof("Starting coordinator.")
	c.logger.Infof("Starting coordinator rpc server.")
	go func() {
		err := c.startServer(ctx)
		if err != nil {
			c.logger.Error("Starting coordinator rpc server error", "err", err)
		}
	}()

	var metricsErrChan <-chan error
	metricsErrChan = c.metrics.Start(ctx, c.metricsReg)

	// subscribe to onchain event
	sub := c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)
	// ticker for creating new task
	ticker := time.NewTicker(time.Duration(c.taskDurationMinutes) * time.Minute)
	c.logger.Infof("Coordinator set to send new task every 2 hours")
	defer ticker.Stop()

	taskDetails, err := c.createNewTask()
	if err != nil {
		c.logger.Error("Create new task error", "err", err)
	} else {
		// ticker doesn't tick immediately, so we send the first task here
		_ = c.sendNewTask(taskDetails)
	}

	for {
		select {
		case err := <-metricsErrChan:
			c.logger.Fatal("Error in metrics server", "err", err)
		case <-ticker.C:
			taskDetails, err = c.createNewTask()
			if err != nil {
				c.logger.Error("Create new task error", "err", err)
				continue
			}
			err := c.sendNewTask(taskDetails)
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to the next task
				continue
			}
		case err := <-sub.Err():
			c.logger.Error("Error in websocket subscription", "err", err)
			sub.Unsubscribe()
			sub = c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)
		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			err := c.handleNewTaskCreatedLog(ctx, newTaskCreatedLog)
			if err != nil {
				continue
			}
		case blsAggServiceResp := <-c.blsAggregationService.GetResponseChannel():
			c.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			c.sendAggregatedResponseToContract(blsAggServiceResp)
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *Coordinator) createNewTask() (string, error) {
	randomIndex := rand.Intn(len(c.taskChains))
	chain := c.taskChains[randomIndex]

	c.logger.Infof("Getting %s latest block height", chain)
	latestBlockHeight, err := c.flinkClient.GetChainLatestBlockHeight(chain)
	if err != nil {
		c.logger.Error("Failed to get latest block height", "chain", chain, "err", err)
		return "", err
	}
	c.logger.Infof("%s latest block height: %d", chain, latestBlockHeight)

	taskDetails := core.GenerateTaskDetails(&core.TaskDetails{
		Version:    "v1",
		Chain:      chain,
		TaskType:   "block",
		Method:     "merkle",
		StartBlock: int(latestBlockHeight),
		EndBlock:   int(latestBlockHeight) + 100,
		Difficulty: 10,
		Deadline:   time.Now().Add(time.Duration(c.taskDurationMinutes) * time.Minute).Unix(),
	})
	return taskDetails, nil
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (c *Coordinator) sendNewTask(taskDetails string) error {
	c.logger.Info("Coordinator sending new task", "task details", taskDetails)
	// Send taskDetails to the task manager contract
	newTask, taskIndex, err := c.avsWriter.SendNewTask(context.Background(), taskDetails, sdktypes.QuorumThresholdPercentage(c.quorumThreshold), types.QuorumNumbers)
	if err != nil {
		c.logger.Error("Coordinator failed to send task", "err", err)
		return err
	}

	c.tasksMu.Lock()
	c.tasks[taskIndex] = newTask
	c.tasksMu.Unlock()

	// metrics
	c.metrics.IncNumTaskCreated()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	parsedTaskDetails, err := core.ParseTaskDetails(taskDetails)
	if err != nil {
		c.logger.Error("Failed to parse task details", "err", err)
		return err
	}
	taskTimeToExpiry := time.Duration(parsedTaskDetails.Deadline-time.Now().Unix()) * time.Second

	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	err = c.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	if err != nil {
		c.logger.Error("blsAggregationService failed to initialize new task", "err", err)
		return err
	}
	return nil
}

func (c *Coordinator) handleNewTaskCreatedLog(ctx context.Context, newTaskCreatedLog *bindings.ChainbaseServiceManagerNewTaskCreated) error {
	curBlockNum, err := c.ethClient.BlockNumber(ctx)
	if err != nil {
		c.logger.Error("Unable to get current block number")
		return err
	}

	quorumNumbers := sdktypes.QuorumNums{sdktypes.QuorumNum(0)}
	operatorsAvsStateDict, err := c.avsRegistryService.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, sdktypes.BlockNum(curBlockNum))
	if err != nil {
		return err
	}
	for _, avsState := range operatorsAvsStateDict {
		c.logger.Info("manuscript node", "OperatorId", hex.EncodeToString(avsState.OperatorId[:]), "Socket", avsState.OperatorInfo.Socket)

		nodeRpcClient, err := NewManuscriptRpcClient(avsState.OperatorInfo.Socket.String(), c.logger, c.metrics)
		if err != nil {
			c.logger.Error("Cannot create ManuscriptRpcClient. Is manuscript node running?", "err", err)
			return err
		}

		nodeRpcClient.CreateNewTask(&nodepb.NewTaskRequest{
			TaskIndex: newTaskCreatedLog.TaskIndex,
			Task: &nodepb.Task{
				TaskDetails:               newTaskCreatedLog.Task.TaskDetails,
				TaskCreatedBlock:          newTaskCreatedLog.Task.TaskCreatedBlock,
				QuorumNumbers:             newTaskCreatedLog.Task.QuorumNumbers,
				QuorumThresholdPercentage: newTaskCreatedLog.Task.QuorumThresholdPercentage,
			},
		})
	}
	c.logger.Info("New task has been successfully sent to all nodes.")
	return nil
}

func (c *Coordinator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	if blsAggServiceResp.Err != nil {
		c.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		return
	}
	if blsAggServiceResp.TaskResponse == nil {
		c.logger.Error("Invalid task response")
		return
	}
	nonSignerPubkeys := []bindings.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []bindings.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := bindings.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	c.logger.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)
	c.tasksMu.RLock()
	task := c.tasks[blsAggServiceResp.TaskIndex]
	c.tasksMu.RUnlock()
	c.taskResponsesMu.RLock()
	taskResponse := c.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	c.taskResponsesMu.RUnlock()
	_, err := c.avsWriter.SendAggregatedResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		c.logger.Error("Coordinator failed to respond to task", "err", err)
	}
	// metrics
	c.metrics.IncNumTaskCompleted()
	c.logger.Info("Aggregated response has been successfully sent to ChainbaseServiceManager contract.")
}
