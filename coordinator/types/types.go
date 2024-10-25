package types

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

var QuorumNumbers = sdktypes.QuorumNums{0}

type BlockNumber = uint32
type TaskIndex = uint32

type OperatorInfo struct {
	OperatorPubkeys sdktypes.OperatorPubkeys
	OperatorAddr    common.Address
}
