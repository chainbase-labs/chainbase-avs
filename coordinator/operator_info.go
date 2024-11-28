package coordinator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

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

	if err := c.updateOperatorsInfo(ctx); err != nil {
		c.logger.Error("Error in update operators info", "err", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := c.updateOperatorsInfo(ctx); err != nil {
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
				TaskID:     newTaskCreatedLog.TaskIndex,
				TaskDetail: newTaskCreatedLog.Task.TaskDetails,
			})
			if err != nil {
				c.logger.Error("Error in insert task", "err", err)
			}
		case taskResponseLog := <-taskResponseChan:
			err := postgres.UpdateTaskResponse(c.db, taskResponseLog.TaskResponse.ReferenceTaskIndex, taskResponseLog.TaskResponse.TaskResponse)
			if err != nil {
				c.logger.Error("Error in update task response", "err", err)
			}
		case <-ctx.Done():
			c.logger.Error("Update routine stopping...")
			return
		}
	}
}

func (c *Coordinator) updateOperatorsInfo(ctx context.Context) error {
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

		nodeRpcClient, err := NewManuscriptRpcClient(operatorInfo.Socket.String(), c.logger, c.metrics)
		if err != nil {
			c.logger.Error("Cannot create ManuscriptRpcClient. Is manuscript node running?", "err", err)
			continue
		}

		operatorAddress := operatorStake.Operator.String()
		operatorHWInfo, err := nodeRpcClient.GetOperatorInfo()
		if err != nil {
			c.logger.Error("Failed to get operator hardware info", "operatorAddress", operatorAddress, "err", err)
			err := postgres.UpdateOperatorStatus(c.db, operatorAddress)
			if err != nil {
				c.logger.Error("Failed to update operator status", "operatorAddress", operatorAddress, "err", err)
			}
			continue
		}

		socket := operatorInfo.Socket.String()
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
			Status:          "active",
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

	operators := make([]common.Address, len(addresses))
	for _, address := range addresses {
		operators = append(operators, common.HexToAddress(address))
	}
	return c.updateOperatorRegisteredAt(operators)
}

func (c *Coordinator) updateOperatorRegisteredAt(operators []common.Address) error {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(c.registryCoordinatorAddr, c.ethClient)
	if err != nil {
		return err
	}

	operatorRegisteredIterator, err := contractRegistryCoordinator.FilterOperatorRegistered(&bind.FilterOpts{Context: context.Background()}, operators, nil)
	if err != nil {
		return err
	}

	for operatorRegisteredIterator.Next() {
		operator := operatorRegisteredIterator.Event.Operator.String()
		blockhash := operatorRegisteredIterator.Event.Raw.BlockHash
		block, err := c.ethClient.BlockByHash(context.Background(), blockhash)
		if err != nil {
			continue
		}

		timestamp := time.Unix(int64(block.Time()), 0)
		err = postgres.UpdateOperatorRegisteredAt(c.db, operator, timestamp)
		if err != nil {
			continue
		}
	}

	return nil
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

	return result.Country, nil
}
