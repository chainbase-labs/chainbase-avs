package coordinator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var ipLocations = make(map[string]string)

func (c *Coordinator) updateOperatorsRoutine(ctx context.Context) {
	// ticker for updating operator info
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	if err := c.updateOperatorsInfo(ctx); err != nil {
		c.logger.Error("Error in update operators info", "err", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := c.updateOperatorsInfo(ctx); err != nil {
				c.logger.Error("Error in update operators info", "err", err)
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
			c.updateOperatorStatus(operatorAddress)
			continue
		}

		socket := operatorInfo.Socket.String()
		location, err := getIpLocation(socket)
		if err != nil {
			return err
		}

		c.saveOperatorInfo(operatorAddress, hex.EncodeToString(operatorStake.OperatorId[:]), socket, location, operatorHWInfo.CpuCore, operatorHWInfo.Memory)
	}
	return c.updateOperatorRegisteredAt()
}

func (c *Coordinator) updateOperatorStatus(operatorAddress string) {
	_, err := c.db.Exec(`UPDATE operator SET status = $1 WHERE operator_address = $2`, "inactive", operatorAddress)
	if err != nil {
		c.logger.Error("Failed to update operator status", "operatorAddress", operatorAddress, "err", err)
	}
}

func (c *Coordinator) saveOperatorInfo(operatorAddress, operatorId, socket, location string, cpuCore, memory uint32) {
	query := `INSERT INTO operator (operator_address,operator_id,socket,location,cpu_core,memory,status,updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, "active", CURRENT_TIMESTAMP)
			ON CONFLICT (operator_address) 
			DO UPDATE SET operator_id = $2,socket = $3,location = $4,cpu_core = $5,memory = $6,status = "active" ,updated_at = CURRENT_TIMESTAMP`

	_, err := c.db.Exec(
		query,
		operatorAddress,
		operatorId,
		socket,
		location,
		cpuCore,
		memory,
		time.Now(),
	)

	if err != nil {
		c.logger.Error("Failed to save operator info", "operator", operatorAddress, "err", err)
	}
}

func (c *Coordinator) updateOperatorRegisteredAt() error {
	rows, err := c.db.Query("SELECT operator_address FROM operator WHERE registered_at IS NULL")
	if err != nil {
		c.logger.Error("Failed to query operators registered at", "err", err)
		return err
	}
	defer rows.Close()

	addresses := make([]string, 0)
	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			c.logger.Error("Failed to scan operator address", "err", err)
			continue
		}
		addresses = append(addresses, address)
	}

	return c.queryOperatorRegisteredAt(addresses)
}

func (c *Coordinator) queryOperatorRegisteredAt(addresses []string) error {
	contractAddress := c.registryCoordinatorAddr

	contractABI := `[{"inputs":[{"internalType":"contract IServiceManager","name":"_serviceManager","type":"address"},{"internalType":"contract IStakeRegistry","name":"_stakeRegistry","type":"address"},{"internalType":"contract IBLSApkRegistry","name":"_blsApkRegistry","type":"address"},{"internalType":"contract IIndexRegistry","name":"_indexRegistry","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"InvalidShortString","type":"error"},{"inputs":[{"internalType":"string","name":"str","type":"string"}],"name":"StringTooLong","type":"error"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"prevChurnApprover","type":"address"},{"indexed":false,"internalType":"address","name":"newChurnApprover","type":"address"}],"name":"ChurnApproverUpdated","type":"event"},{"anonymous":false,"inputs":[],"name":"EIP712DomainChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"prevEjector","type":"address"},{"indexed":false,"internalType":"address","name":"newEjector","type":"address"}],"name":"EjectorUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"version","type":"uint8"}],"name":"Initialized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"bytes32","name":"operatorId","type":"bytes32"}],"name":"OperatorDeregistered","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"bytes32","name":"operatorId","type":"bytes32"}],"name":"OperatorRegistered","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint8","name":"quorumNumber","type":"uint8"},{"components":[{"internalType":"uint32","name":"maxOperatorCount","type":"uint32"},{"internalType":"uint16","name":"kickBIPsOfOperatorStake","type":"uint16"},{"internalType":"uint16","name":"kickBIPsOfTotalStake","type":"uint16"}],"indexed":false,"internalType":"struct IRegistryCoordinator.OperatorSetParam","name":"operatorSetParams","type":"tuple"}],"name":"OperatorSetParamsUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"indexed":false,"internalType":"string","name":"socket","type":"string"}],"name":"OperatorSocketUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":false,"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"Paused","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"contract IPauserRegistry","name":"pauserRegistry","type":"address"},{"indexed":false,"internalType":"contract IPauserRegistry","name":"newPauserRegistry","type":"address"}],"name":"PauserRegistrySet","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint8","name":"quorumNumber","type":"uint8"},{"indexed":false,"internalType":"uint256","name":"blocknumber","type":"uint256"}],"name":"QuorumBlockNumberUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":false,"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"Unpaused","type":"event"},{"inputs":[],"name":"OPERATOR_CHURN_APPROVAL_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"PUBKEY_REGISTRATION_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"blsApkRegistry","outputs":[{"internalType":"contract IBLSApkRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"registeringOperator","type":"address"},{"internalType":"bytes32","name":"registeringOperatorId","type":"bytes32"},{"components":[{"internalType":"uint8","name":"quorumNumber","type":"uint8"},{"internalType":"address","name":"operator","type":"address"}],"internalType":"struct IRegistryCoordinator.OperatorKickParam[]","name":"operatorKickParams","type":"tuple[]"},{"internalType":"bytes32","name":"salt","type":"bytes32"},{"internalType":"uint256","name":"expiry","type":"uint256"}],"name":"calculateOperatorChurnApprovalDigestHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"churnApprover","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"components":[{"internalType":"uint32","name":"maxOperatorCount","type":"uint32"},{"internalType":"uint16","name":"kickBIPsOfOperatorStake","type":"uint16"},{"internalType":"uint16","name":"kickBIPsOfTotalStake","type":"uint16"}],"internalType":"struct IRegistryCoordinator.OperatorSetParam","name":"operatorSetParams","type":"tuple"},{"internalType":"uint96","name":"minimumStake","type":"uint96"},{"components":[{"internalType":"contract IStrategy","name":"strategy","type":"address"},{"internalType":"uint96","name":"multiplier","type":"uint96"}],"internalType":"struct IStakeRegistry.StrategyParams[]","name":"strategyParams","type":"tuple[]"}],"name":"createQuorum","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"quorumNumbers","type":"bytes"}],"name":"deregisterOperator","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"eip712Domain","outputs":[{"internalType":"bytes1","name":"fields","type":"bytes1"},{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"version","type":"string"},{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"address","name":"verifyingContract","type":"address"},{"internalType":"bytes32","name":"salt","type":"bytes32"},{"internalType":"uint256[]","name":"extensions","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"}],"name":"ejectOperator","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"ejector","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"}],"name":"getCurrentQuorumBitmap","outputs":[{"internalType":"uint192","name":"","type":"uint192"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"getOperator","outputs":[{"components":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"enum IRegistryCoordinator.OperatorStatus","name":"status","type":"uint8"}],"internalType":"struct IRegistryCoordinator.OperatorInfo","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"}],"name":"getOperatorFromId","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"getOperatorId","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint8","name":"quorumNumber","type":"uint8"}],"name":"getOperatorSetParams","outputs":[{"components":[{"internalType":"uint32","name":"maxOperatorCount","type":"uint32"},{"internalType":"uint16","name":"kickBIPsOfOperatorStake","type":"uint16"},{"internalType":"uint16","name":"kickBIPsOfTotalStake","type":"uint16"}],"internalType":"struct IRegistryCoordinator.OperatorSetParam","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"getOperatorStatus","outputs":[{"internalType":"enum IRegistryCoordinator.OperatorStatus","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"uint32","name":"blockNumber","type":"uint32"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getQuorumBitmapAtBlockNumberByIndex","outputs":[{"internalType":"uint192","name":"","type":"uint192"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"}],"name":"getQuorumBitmapHistoryLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint32","name":"blockNumber","type":"uint32"},{"internalType":"bytes32[]","name":"operatorIds","type":"bytes32[]"}],"name":"getQuorumBitmapIndicesAtBlockNumber","outputs":[{"internalType":"uint32[]","name":"","type":"uint32[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getQuorumBitmapUpdateByIndex","outputs":[{"components":[{"internalType":"uint32","name":"updateBlockNumber","type":"uint32"},{"internalType":"uint32","name":"nextUpdateBlockNumber","type":"uint32"},{"internalType":"uint192","name":"quorumBitmap","type":"uint192"}],"internalType":"struct IRegistryCoordinator.QuorumBitmapUpdate","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"indexRegistry","outputs":[{"internalType":"contract IIndexRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_initialOwner","type":"address"},{"internalType":"address","name":"_churnApprover","type":"address"},{"internalType":"address","name":"_ejector","type":"address"},{"internalType":"contract IPauserRegistry","name":"_pauserRegistry","type":"address"},{"internalType":"uint256","name":"_initialPausedStatus","type":"uint256"},{"components":[{"internalType":"uint32","name":"maxOperatorCount","type":"uint32"},{"internalType":"uint16","name":"kickBIPsOfOperatorStake","type":"uint16"},{"internalType":"uint16","name":"kickBIPsOfTotalStake","type":"uint16"}],"internalType":"struct IRegistryCoordinator.OperatorSetParam[]","name":"_operatorSetParams","type":"tuple[]"},{"internalType":"uint96[]","name":"_minimumStakes","type":"uint96[]"},{"components":[{"internalType":"contract IStrategy","name":"strategy","type":"address"},{"internalType":"uint96","name":"multiplier","type":"uint96"}],"internalType":"struct IStakeRegistry.StrategyParams[][]","name":"_strategyParams","type":"tuple[][]"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"isChurnApproverSaltUsed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"numRegistries","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"pauseAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint8","name":"index","type":"uint8"}],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pauserRegistry","outputs":[{"internalType":"contract IPauserRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"pubkeyRegistrationMessageHash","outputs":[{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"quorumCount","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint8","name":"","type":"uint8"}],"name":"quorumUpdateBlockNumber","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"string","name":"socket","type":"string"},{"components":[{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"pubkeyRegistrationSignature","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"pubkeyG1","type":"tuple"},{"components":[{"internalType":"uint256[2]","name":"X","type":"uint256[2]"},{"internalType":"uint256[2]","name":"Y","type":"uint256[2]"}],"internalType":"struct BN254.G2Point","name":"pubkeyG2","type":"tuple"}],"internalType":"struct IBLSApkRegistry.PubkeyRegistrationParams","name":"params","type":"tuple"},{"components":[{"internalType":"bytes","name":"signature","type":"bytes"},{"internalType":"bytes32","name":"salt","type":"bytes32"},{"internalType":"uint256","name":"expiry","type":"uint256"}],"internalType":"struct ISignatureUtils.SignatureWithSaltAndExpiry","name":"operatorSignature","type":"tuple"}],"name":"registerOperator","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"string","name":"socket","type":"string"},{"components":[{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"pubkeyRegistrationSignature","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"pubkeyG1","type":"tuple"},{"components":[{"internalType":"uint256[2]","name":"X","type":"uint256[2]"},{"internalType":"uint256[2]","name":"Y","type":"uint256[2]"}],"internalType":"struct BN254.G2Point","name":"pubkeyG2","type":"tuple"}],"internalType":"struct IBLSApkRegistry.PubkeyRegistrationParams","name":"params","type":"tuple"},{"components":[{"internalType":"uint8","name":"quorumNumber","type":"uint8"},{"internalType":"address","name":"operator","type":"address"}],"internalType":"struct IRegistryCoordinator.OperatorKickParam[]","name":"operatorKickParams","type":"tuple[]"},{"components":[{"internalType":"bytes","name":"signature","type":"bytes"},{"internalType":"bytes32","name":"salt","type":"bytes32"},{"internalType":"uint256","name":"expiry","type":"uint256"}],"internalType":"struct ISignatureUtils.SignatureWithSaltAndExpiry","name":"churnApproverSignature","type":"tuple"},{"components":[{"internalType":"bytes","name":"signature","type":"bytes"},{"internalType":"bytes32","name":"salt","type":"bytes32"},{"internalType":"uint256","name":"expiry","type":"uint256"}],"internalType":"struct ISignatureUtils.SignatureWithSaltAndExpiry","name":"operatorSignature","type":"tuple"}],"name":"registerOperatorWithChurn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"registries","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"serviceManager","outputs":[{"internalType":"contract IServiceManager","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_churnApprover","type":"address"}],"name":"setChurnApprover","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_ejector","type":"address"}],"name":"setEjector","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint8","name":"quorumNumber","type":"uint8"},{"components":[{"internalType":"uint32","name":"maxOperatorCount","type":"uint32"},{"internalType":"uint16","name":"kickBIPsOfOperatorStake","type":"uint16"},{"internalType":"uint16","name":"kickBIPsOfTotalStake","type":"uint16"}],"internalType":"struct IRegistryCoordinator.OperatorSetParam","name":"operatorSetParams","type":"tuple"}],"name":"setOperatorSetParams","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"contract IPauserRegistry","name":"newPauserRegistry","type":"address"}],"name":"setPauserRegistry","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"stakeRegistry","outputs":[{"internalType":"contract IStakeRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[]","name":"operators","type":"address[]"}],"name":"updateOperators","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[][]","name":"operatorsPerQuorum","type":"address[][]"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"}],"name":"updateOperatorsForQuorum","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"socket","type":"string"}],"name":"updateSocket","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return err
	}

	methodName := "registerOperator"
	methodID := parsedABI.Methods[methodName].ID

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
	}

	logs, err := c.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		c.logger.Error("Failed to get logs", "err", err)
		return err
	}

	for _, vLog := range logs {
		if len(vLog.Data) >= 4 && strings.HasPrefix(common.Bytes2Hex(vLog.Data), common.Bytes2Hex(methodID)) {
			tx, _, err := c.ethClient.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				continue
			}

			from, err := c.ethClient.TransactionSender(context.Background(), tx, vLog.BlockHash, vLog.TxIndex)
			if err != nil {
				continue
			}

			for _, address := range addresses {
				if from == common.HexToAddress(address) {
					block, err := c.ethClient.BlockByHash(context.Background(), vLog.BlockHash)
					if err != nil {
						continue
					}

					timestamp := time.Unix(int64(block.Time()), 0)
					_, err = c.db.Exec(`UPDATE operator SET registered_at = $1 WHERE operator_address = $2`, timestamp, address)
					if err != nil {
						continue
					}
				}
			}
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
