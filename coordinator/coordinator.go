package coordinator

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/coordinator/types"
	"github.com/chainbase-labs/chainbase-avs/core"
	"github.com/chainbase-labs/chainbase-avs/core/chainio"
	"github.com/chainbase-labs/chainbase-avs/core/config"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTime                = 12 * time.Second
	avsName                  = "Chainbase"
)

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
	avsWriter        chainio.IAvsWriter
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskIndex]bindings.IChainbaseServiceManagerTask
	tasksMu               sync.RWMutex
	taskResponses         map[types.TaskIndex]map[sdktypes.TaskResponseDigest]bindings.IChainbaseServiceManagerTaskResponse
	taskResponsesMu       sync.RWMutex
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

	chainConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	chainClients, err := clients.BuildAll(chainConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create sdk clients", "err", err)
		return nil, err
	}

	hashFunction := func(taskResponse sdktypes.TaskResponse) (sdktypes.TaskResponseDigest, error) {
		taskResponseBytes, err := json.Marshal(taskResponse)
		if err != nil {
			return sdktypes.TaskResponseDigest{}, err
		}

		return sdktypes.TaskResponseDigest(taskResponseBytes), nil
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), chainClients.AvsRegistryChainSubscriber, chainClients.AvsRegistryChainReader, nil, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, hashFunction, c.Logger)

	return &Coordinator{
		logger:                c.Logger,
		serverIpPortAddr:      c.CoordinatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskIndex]bindings.IChainbaseServiceManagerTask),
		taskResponses:         make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]bindings.IChainbaseServiceManagerTaskResponse),
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

	ticker := time.NewTicker(2 * time.Hour)
	c.logger.Infof("Coordinator set to send new task every 2 hours")
	defer ticker.Stop()
	// TODO
	taskDetails := ""
	// ticker doesn't tick immediately, so we send the first task here
	_ = c.sendNewTask(taskDetails)

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-c.blsAggregationService.GetResponseChannel():
			c.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			c.sendAggregatedResponseToContract(blsAggServiceResp)
		case <-ticker.C:
			err := c.sendNewTask(taskDetails)
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to the next task
				continue
			}
		}
	}
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (c *Coordinator) sendNewTask(taskDetails string) error {
	c.logger.Info("Coordinator sending new task", "task details", taskDetails)
	// Send taskDetails to the task manager contract
	newTask, taskIndex, err := c.avsWriter.SendNewTask(context.Background(), taskDetails, types.QuorumThresholdNumerator, types.QuorumNumbers)
	if err != nil {
		c.logger.Error("Coordinator failed to send task", "err", err)
		return err
	}

	c.tasksMu.Lock()
	c.tasks[taskIndex] = newTask
	c.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	//TODO parse taskTimeToExpiry from task detail
	taskTimeToExpiry := taskChallengeWindowBlock * blockTime
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

func (c *Coordinator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		c.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panice to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
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
}
