package coordinator

import (
	"context"
	"encoding/hex"
	"errors"
	"net"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"google.golang.org/grpc"

	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/core"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (c *Coordinator) startServer(_ context.Context) error {
	listener, err := net.Listen("tcp", c.serverIpPortAddr)
	if err != nil {
		c.logger.Fatal("Failed to listen", "err", err)
	}

	server := grpc.NewServer()
	coordinatorpb.RegisterCoordinatorServiceServer(server, c)
	c.logger.Info("Server listening at", "addr", listener.Addr())

	if err := server.Serve(listener); err != nil {
		c.logger.Fatal("Failed to start serve", "err", err)
	}
	return nil
}

// ProcessSignedTaskResponse rpc endpoint which is called by manuscript node
func (c *Coordinator) ProcessSignedTaskResponse(ctx context.Context, signedTaskResponse *coordinatorpb.SignedTaskResponseRequest) (*coordinatorpb.SignedTaskResponseReply, error) {
	c.logger.Infof("Received signed task response from operator %s", hex.EncodeToString(signedTaskResponse.OperatorId))

	taskResponse := bindings.IChainbaseServiceManagerTaskResponse{
		ReferenceTaskIndex: signedTaskResponse.TaskResponse.ReferenceTaskIndex,
		TaskResponse:       signedTaskResponse.TaskResponse.TaskResponse,
	}

	taskIndex := taskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetTaskResponseDigest(&taskResponse)
	if err != nil {
		c.logger.Error("Failed to get task response digest", "err", err)
		return nil, TaskResponseDigestNotFoundError500
	}

	c.taskResponsesMu.Lock()
	if _, ok := c.taskResponses[taskIndex]; !ok {
		c.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]bindings.IChainbaseServiceManagerTaskResponse)
	}
	if _, ok := c.taskResponses[taskIndex][taskResponseDigest]; !ok {
		c.taskResponses[taskIndex][taskResponseDigest] = taskResponse
	}
	c.taskResponsesMu.Unlock()

	// Convert protobuf Signature to bls.Signature
	g1Point := &bn254.G1Affine{
		X: fp.Element{signedTaskResponse.BlsSignature.G1Point.X[0], signedTaskResponse.BlsSignature.G1Point.X[1], signedTaskResponse.BlsSignature.G1Point.X[2], signedTaskResponse.BlsSignature.G1Point.X[3]},
		Y: fp.Element{signedTaskResponse.BlsSignature.G1Point.Y[0], signedTaskResponse.BlsSignature.G1Point.Y[1], signedTaskResponse.BlsSignature.G1Point.Y[2], signedTaskResponse.BlsSignature.G1Point.Y[3]},
	}
	blsSignature := bls.Signature{G1Point: &bls.G1Point{G1Affine: g1Point}}

	var operatorId types.OperatorId
	copy(operatorId[:], signedTaskResponse.OperatorId)

	err = c.blsAggregationService.ProcessNewSignature(
		ctx, taskIndex, taskResponse,
		&blsSignature, operatorId,
	)

	if err != nil {
		c.logger.Error("Failed to process new signature", "err", err)
		return nil, UnknownErrorWhileVerifyingSignature400
	}

	return &coordinatorpb.SignedTaskResponseReply{Success: true}, nil
}
