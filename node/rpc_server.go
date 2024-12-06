package node

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/node/metrics"
)

func (n *ManuscriptNode) startServer(_ context.Context) error {
	listener, err := net.Listen("tcp", n.nodeGrpcServerAddress)
	if err != nil {
		n.logger.Fatal("Failed to listen", "err", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(n.ipCheckInterceptor),
	)
	nodepb.RegisterManuscriptNodeServiceServer(server, n)
	n.logger.Info("Server listening at", "addr", listener.Addr())

	if err := server.Serve(listener); err != nil {
		n.logger.Fatal("Failed to start serve", "err", err)
	}
	return nil
}

func (n *ManuscriptNode) ipCheckInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get peer info")
	}

	addr := p.Addr.String()
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to parse client ip")
	}

	allowedIP, _, err := net.SplitHostPort(n.coordinatorServerIpPortAddr)
	if err != nil {
		n.logger.Error("Failed to parse coordinator ip", "err", err)
		return nil, status.Errorf(codes.Internal, "failed to parse coordinator ip")
	}

	if ip != allowedIP {
		n.logger.Info("Client ip is not allowed", "ip", ip)
		return nil, status.Errorf(codes.PermissionDenied, "client ip is not allowed")
	}

	return handler(ctx, req)
}

func (n *ManuscriptNode) ReceiveNewTask(_ context.Context, req *nodepb.NewTaskRequest) (*nodepb.NewTaskResponse, error) {
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

func (n *ManuscriptNode) GetOperatorInfo(_ context.Context, _ *nodepb.GetOperatorInfoRequest) (*nodepb.GetOperatorInfoResponse, error) {
	memory := metrics.GetTotalMemory()
	return &nodepb.GetOperatorInfoResponse{
		CpuCore: metrics.GetCPUCore(),
		Memory:  uint32(memory),
	}, nil
}
