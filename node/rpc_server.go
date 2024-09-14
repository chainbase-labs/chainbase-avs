package node

import (
	"context"
	"net"

	"google.golang.org/grpc"

	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
)

func (n *ManuscriptNode) startServer(_ context.Context) error {
	listener, err := net.Listen("tcp", n.nodeServerIpPortAddr)
	if err != nil {
		n.logger.Fatal("Failed to listen", "err", err)
	}

	server := grpc.NewServer()
	nodepb.RegisterManuscriptNodeServiceServer(server, n)
	n.logger.Info("Server listening at", "addr", listener.Addr())

	if err := server.Serve(listener); err != nil {
		n.logger.Fatal("Failed to start serve", "err", err)
	}
	return nil
}

func (n *ManuscriptNode) ReceiveNewTask(_ context.Context, req *nodepb.NewTaskRequest) (*nodepb.NewTaskResponse, error) {
	// TODO verify request from coordinator
	newTask := &bindings.ChainbaseServiceManagerNewTaskCreated{
		TaskIndex: req.TaskIndex,
		Task: bindings.IChainbaseServiceManagerTask{
			TaskDetails:               req.Task.TaskDetails,
			TaskCreatedBlock:          req.Task.TaskCreatedBlock,
			QuorumNumbers:             req.Task.QuorumNumbers,
			QuorumThresholdPercentage: req.Task.QuorumThresholdPercentage,
		},
	}

	n.newTaskCreatedChan <- newTask

	return &nodepb.NewTaskResponse{
		Success: true,
	}, nil
}
