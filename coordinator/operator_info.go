package coordinator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/coordinator/postgres"
)

var ipLocations = make(map[string]string)

func (c *Coordinator) updateOperatorsRoutine(ctx context.Context) {
	// ticker for updating operator info
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	newTaskCreatedChan := make(chan *bindings.ChainbaseServiceManagerNewTaskCreated)
	newTaskCreatedSub := c.avsSubscriber.SubscribeToNewTasks(newTaskCreatedChan)

	taskResponseChan := make(chan *bindings.ChainbaseServiceManagerTaskResponded)
	taskResponseSub := c.avsSubscriber.SubscribeToTaskResponses(taskResponseChan)

	//update exist operator info
	if err := c.updateOperators(ctx); err != nil {
		c.logger.Error("Error in update operators", "err", err)
	}

	// update exist tasks
	if err := c.updateTasks(ctx); err != nil {
		c.logger.Error("Error in update tasks", "err", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := c.updateOperators(ctx); err != nil {
				c.logger.Error("Error in update operators info", "err", err)
			}
		case err := <-newTaskCreatedSub.Err():
			c.logger.Error("Error in new task created subscription", "err", err)
			newTaskCreatedSub.Unsubscribe()
			newTaskCreatedSub = c.avsSubscriber.SubscribeToTaskResponses(taskResponseChan)
		case err := <-taskResponseSub.Err():
			c.logger.Error("Error in task response subscription", "err", err)
			taskResponseSub.Unsubscribe()
			taskResponseSub = c.avsSubscriber.SubscribeToTaskResponses(taskResponseChan)
		case newTaskCreatedLog := <-newTaskCreatedChan:
			err := postgres.InsertTask(c.db, &postgres.Task{
				TaskID:       newTaskCreatedLog.TaskIndex,
				TaskDetail:   newTaskCreatedLog.Task.TaskDetails,
				CreateTaskTx: newTaskCreatedLog.Raw.TxHash.String(),
			})
			if err != nil {
				c.logger.Error("Error in insert task", "err", err)
			}
		case taskResponseLog := <-taskResponseChan:
			err := postgres.UpdateTaskResponse(c.db,
				taskResponseLog.TaskResponse.ReferenceTaskIndex,
				taskResponseLog.TaskResponse.TaskResponse,
				taskResponseLog.Raw.TxHash.String(),
			)
			if err != nil {
				c.logger.Error("Error in update task response", "err", err)
			}
			err = c.updateOperatorTasks(ctx,
				taskResponseLog.TaskResponse.ReferenceTaskIndex,
				taskResponseLog.Raw.TxHash,
				uint32(taskResponseLog.Raw.BlockNumber),
			)
			if err != nil {
				c.logger.Error("Error in update operator tasks", "err", err)
			}
		case <-ctx.Done():
			c.logger.Error("Update routine stopping...")
			return
		}
	}
}

func (c *Coordinator) updateOperators(ctx context.Context) error {
	quorumNumbers := sdktypes.QuorumNums{sdktypes.QuorumNum(0)}
	// Get operator state for each quorum by querying OperatorStateRetriever
	operatorsStakesInQuorums, err := c.avsRegistryService.GetOperatorsStakeInQuorumsAtCurrentBlock(&bind.CallOpts{Context: ctx}, quorumNumbers)
	if err != nil {
		return errors.Wrap(err, "Failed to get operator state")
	}
	if len(operatorsStakesInQuorums) != len(quorumNumbers) {
		return errors.Wrap(err, "Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested")
	}

	operatorsStakes := operatorsStakesInQuorums[0]
	for _, operatorStake := range operatorsStakes {
		operatorInfo, exist := c.operatorsInfoService.GetOperatorInfo(ctx, operatorStake.Operator)
		if !exist {
			c.logger.Error("Failed to find operator info")
			continue
		}

		socket := operatorInfo.Socket.String()
		nodeRpcClient, err := NewManuscriptRpcClient(socket, c.logger, c.metrics)
		if err != nil {
			c.logger.Error("Cannot create ManuscriptRpcClient. Is manuscript node running?", "err", err)
			continue
		}

		operatorStatus := "active"
		operatorAddress := operatorStake.Operator.String()
		operatorHWInfo, err := nodeRpcClient.GetOperatorInfo()
		if err != nil {
			c.logger.Error("Failed to get operator hardware info", "operatorAddress", operatorAddress, "err", err)
			operatorHWInfo = &nodepb.GetOperatorInfoResponse{CpuCore: 0, Memory: 0}
			operatorStatus = "inactive"
		}

		location, err := getIpLocation(socket)
		if err != nil {
			c.logger.Error("Failed to get operator location", "socket", socket, "err", err)
			continue
		}

		_, err = postgres.UpsertOperator(c.db, &postgres.Operator{
			OperatorAddress: operatorAddress,
			OperatorID:      hex.EncodeToString(operatorStake.OperatorId[:]),
			Socket:          socket,
			Location:        location,
			CPUCore:         operatorHWInfo.CpuCore,
			Memory:          operatorHWInfo.Memory,
			Status:          operatorStatus,
		})
		if err != nil {
			c.logger.Error("Failed to upsert operator", "operatorAddress", operatorAddress, "err", err)
			continue
		}
	}

	addresses, err := postgres.QueryOperatorAddressesNoRegisteredAt(c.db)
	if err != nil {
		c.logger.Error("Failed to query operators no registered at", "err", err)
	}

	operators := make([]common.Address, 0)
	for _, address := range addresses {
		operators = append(operators, common.HexToAddress(address))
	}
	return c.updateOperatorRegisteredAt(ctx, operators)
}

func (c *Coordinator) updateOperatorRegisteredAt(ctx context.Context, operators []common.Address) error {
	batchSize := 500
	for len(operators) > 0 {
		end := len(operators)
		if len(operators) > batchSize {
			end = batchSize
		}

		operatorRegisteredIterator, err := c.registryCoordinator.FilterOperatorRegistered(&bind.FilterOpts{Context: ctx}, operators[:end], nil)
		if err != nil {
			return err
		}

		for operatorRegisteredIterator.Next() {
			operator := operatorRegisteredIterator.Event.Operator.String()
			blockhash := operatorRegisteredIterator.Event.Raw.BlockHash
			block, err := c.ethClient.BlockByHash(ctx, blockhash)
			if err != nil {
				continue
			}

			timestamp := time.Unix(int64(block.Time()), 0)
			err = postgres.UpdateOperatorRegisteredAt(c.db, operator, timestamp)
			if err != nil {
				continue
			}
		}

		operators = operators[end:]
	}

	return nil
}

func (c *Coordinator) updateTasks(ctx context.Context) error {
	taskCreatedIterator, err := c.avsReader.AvsServiceBindings.ServiceManager.FilterNewTaskCreated(&bind.FilterOpts{Context: ctx}, nil)
	if err != nil {
		c.logger.Error("Error in filter new task", "err", err)
		return err
	}

	for taskCreatedIterator.Next() {
		err := postgres.InsertTask(c.db, &postgres.Task{
			TaskID:       taskCreatedIterator.Event.TaskIndex,
			TaskDetail:   taskCreatedIterator.Event.Task.TaskDetails,
			CreateTaskTx: taskCreatedIterator.Event.Raw.TxHash.String(),
		})
		if err != nil {
			c.logger.Error("Error in insert task", "err", err)
			return err
		}
	}

	taskResponseIterator, err := c.avsReader.AvsServiceBindings.ServiceManager.FilterTaskResponded(&bind.FilterOpts{Context: ctx})
	if err != nil {
		c.logger.Error("Error in filter task response", "err", err)
		return err
	}

	for taskResponseIterator.Next() {
		err := postgres.UpdateTaskResponse(c.db,
			taskResponseIterator.Event.TaskResponse.ReferenceTaskIndex,
			taskResponseIterator.Event.TaskResponse.TaskResponse,
			taskResponseIterator.Event.Raw.TxHash.String(),
		)
		if err != nil {
			c.logger.Error("Error in update task response", "err", err)
			return err
		}

		err = c.updateOperatorTasks(ctx,
			taskResponseIterator.Event.TaskResponse.ReferenceTaskIndex,
			taskResponseIterator.Event.Raw.TxHash,
			uint32(taskResponseIterator.Event.Raw.BlockNumber),
		)
		if err != nil {
			c.logger.Error("Error in update operator tasks", "err", err)
			return err
		}
	}

	return nil
}

type ResponseTxInput struct {
	Task                        bindings.IChainbaseServiceManagerTask
	TaskResponse                bindings.IChainbaseServiceManagerTaskResponse
	NonSignerStakesAndSignature bindings.IBLSSignatureCheckerNonSignerStakesAndSignature
}

func (c *Coordinator) updateOperatorTasks(ctx context.Context, taskID uint32, txHash common.Hash, blockNum uint32) error {
	operatorTaskCount, err := postgres.CountOperatorTasks(c.db, taskID)
	if err != nil {
		return errors.Wrap(err, "failed to count operator tasks")
	}

	// skip if operator task is already processed
	if operatorTaskCount != 0 {
		return nil
	}

	tx, _, err := c.ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	parsedABI, err := abi.JSON(strings.NewReader(bindings.ChainbaseServiceManagerMetaData.ABI))
	if err != nil {
		return err
	}

	data := tx.Data()
	if len(data) < 4 {
		return errors.New("invalid data for unpacking")
	}

	method, err := parsedABI.MethodById(data[:4])
	if err != nil {
		return err
	}

	var input ResponseTxInput
	err = UnpackIntoInterface(parsedABI, &input, method.Name, data[4:])
	if err != nil {
		return err
	}

	nonSignerPubkeys := input.NonSignerStakesAndSignature.NonSignerPubkeys
	nonSignerOperatorIds := make(map[string]bool, len(nonSignerPubkeys))
	for _, pubkey := range nonSignerPubkeys {
		operatorId := sdktypes.OperatorIdFromContractG1Pubkey(apkreg.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		})
		nonSignerOperatorIds[hex.EncodeToString(operatorId[:])] = true
	}

	quorumNumbers := sdktypes.QuorumNums{sdktypes.QuorumNum(0)}
	operatorsStakesInQuorums, err := c.avsRegistryService.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, blockNum)
	if err != nil {
		return errors.Wrap(err, "Failed to get operator state")
	}
	if len(operatorsStakesInQuorums) != len(quorumNumbers) {
		return errors.Wrap(err, "Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested")
	}
	operatorsStakes := operatorsStakesInQuorums[0]

	signerAddresses := make([]string, 0)
	nonSignerAddresses := make([]string, 0)
	for _, operatorStake := range operatorsStakes {
		operatorId := hex.EncodeToString(operatorStake.OperatorId[:])
		if nonSignerOperatorIds[operatorId] {
			nonSignerAddresses = append(nonSignerAddresses, operatorStake.Operator.String())
		} else {
			signerAddresses = append(signerAddresses, operatorStake.Operator.String())
		}
	}

	if err = c.InsertOperatorTasks(err, signerAddresses, taskID, "completed"); err != nil {
		c.logger.Error("Failed to insert completed operator tasks", "err", err)
	}

	if err = c.InsertOperatorTasks(err, nonSignerAddresses, taskID, "failed"); err != nil {
		c.logger.Error("Failed to insert failed operator tasks", "err", err)
	}

	return nil
}

func (c *Coordinator) InsertOperatorTasks(err error, addresses []string, taskID uint32, status string) error {
	if len(addresses) == 0 {
		return nil
	}

	operatorIDs, err := postgres.QueryOperatorIDs(c.db, addresses)
	if err != nil {
		return err
	}

	if len(operatorIDs) == 0 {
		return nil
	}

	signerOperatorTasks := make([]*postgres.OperatorTask, 0)
	for _, operatorID := range operatorIDs {
		signerOperatorTasks = append(signerOperatorTasks, &postgres.OperatorTask{
			OperatorID: operatorID,
			TaskID:     taskID,
			Status:     status,
		})
	}

	return postgres.BatchInsertOperatorTasks(c.db, signerOperatorTasks)
}

type IpApiResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func getIpLocation(socket string) (string, error) {
	ip := strings.Split(socket, ":")[0]

	if location, ok := ipLocations[ip]; ok {
		return location, nil
	}

	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result IpApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Status != "success" {
		return "", fmt.Errorf("failed to get location for IP %s", ip)
	}

	ipLocations[ip] = result.Country

	time.Sleep(time.Second)

	return result.Country, nil
}

func UnpackIntoInterface(parsedABI abi.ABI, v interface{}, name string, data []byte) error {
	var args abi.Arguments
	if method, ok := parsedABI.Methods[name]; ok {
		if len(data)%32 != 0 {
			return errors.New("abi: improperly formatted output")
		}
		args = method.Inputs
	}
	if args == nil {
		return errors.New("abi: could not locate named method")
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}
