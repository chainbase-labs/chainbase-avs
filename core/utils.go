package core

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
)

func AbiEncodeTaskResponse(h *bindings.IChainbaseServiceManagerTaskResponse) ([]byte, error) {
	// The order here has to match the field ordering of bindings.IChainbaseServiceManagerTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "taskResponse",
			Type: "string",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetTaskResponseDigest(h *bindings.IChainbaseServiceManagerTaskResponse) ([32]byte, error) {
	encodeTaskResponseByte, err := AbiEncodeTaskResponse(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

// BINDING UTILS - conversion from contract structs to golang structs

// ConvertToBN254G1Point BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) bindings.BN254G1Point {
	output := bindings.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) bindings.BN254G2Point {
	output := bindings.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}

type TaskDetails struct {
	Version    string
	Chain      string
	TaskType   string
	Method     string
	StartBlock int
	EndBlock   int
	Difficulty int
	Deadline   int64
}

func GenerateTaskDetails(task *TaskDetails) string {
	parts := []string{
		task.Version,
		task.Chain,
		task.TaskType,
		task.Method,
		"start:" + strconv.Itoa(task.StartBlock),
		"end:" + strconv.Itoa(task.EndBlock),
		"difficulty:" + strconv.Itoa(task.Difficulty),
		"deadline:" + strconv.FormatInt(task.Deadline, 10),
	}
	return strings.Join(parts, ";")
}

func ParseTaskDetails(details string) (*TaskDetails, error) {
	parts := strings.Split(details, ";")
	if len(parts) < 8 {
		return nil, fmt.Errorf("invalid task details format")
	}

	startBlock, err := strconv.Atoi(strings.Split(parts[4], ":")[1])
	if err != nil {
		return nil, err
	}

	endBlock, err := strconv.Atoi(strings.Split(parts[5], ":")[1])
	if err != nil {
		return nil, err
	}

	difficulty, err := strconv.Atoi(strings.Split(parts[6], ":")[1])
	if err != nil {
		return nil, err
	}

	deadline, err := strconv.ParseInt(strings.Split(parts[7], ":")[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &TaskDetails{
		Version:    parts[0],
		Chain:      parts[1],
		TaskType:   parts[2],
		Method:     parts[3],
		StartBlock: startBlock,
		EndBlock:   endBlock,
		Difficulty: difficulty,
		Deadline:   deadline,
	}, nil
}
