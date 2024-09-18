package node

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	coordinatorpb "github.com/chainbase-labs/chainbase-avs/api/grpc/coordinator"
	"github.com/chainbase-labs/chainbase-avs/metrics"
)

type CoordinatorRpcClienter interface {
	SendSignedTaskResponseToCoordinator(signedTaskResponse *coordinatorpb.SignedTaskResponseRequest)
}
type CoordinatorRpcClient struct {
	rpcClient             coordinatorpb.CoordinatorServiceClient
	metrics               metrics.Metrics
	logger                logging.Logger
	coordinatorIpPortAddr string
}

func NewCoordinatorRpcClient(coordinatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*CoordinatorRpcClient, error) {
	return &CoordinatorRpcClient{
		// set to nil so that we can create an rpc client even if the coordinator is not running
		rpcClient:             nil,
		metrics:               metrics,
		logger:                logger,
		coordinatorIpPortAddr: coordinatorIpPortAddr,
	}, nil
}

func (c *CoordinatorRpcClient) dialCoordinatorRpcClient() error {
	client, err := grpc.NewClient(c.coordinatorIpPortAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.rpcClient = coordinatorpb.NewCoordinatorServiceClient(client)
	return nil
}

// SendSignedTaskResponseToCoordinator sends a signed task response to the coordinator.
// it is meant to be ran inside a go thread, so doesn't return anything.
// this is because sending the signed task response to the coordinator is time sensitive,
// so there is no point in retrying if it fails for a few times.
// Currently hardcoded to retry sending the signed task response 5 times, waiting 2 seconds in between each attempt.
func (c *CoordinatorRpcClient) SendSignedTaskResponseToCoordinator(signedTaskResponse *coordinatorpb.SignedTaskResponseRequest) {
	if c.rpcClient == nil {
		c.logger.Info("rpc client is nil. Dialing coordinator rpc client")
		err := c.dialCoordinatorRpcClient()
		if err != nil {
			c.logger.Error("Could not dial coordinator rpc client. Not sending signed task response header to coordinator. Is coordinator running?", "err", err)
			return
		}
	}

	c.logger.Info("Sending signed task response to coordinator", "task index", signedTaskResponse.TaskResponse.ReferenceTaskIndex, "task response", signedTaskResponse.TaskResponse.TaskResponse)
	for i := 0; i < 5; i++ {
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		//defer cancel()

		response, err := c.rpcClient.ProcessSignedTaskResponse(context.Background(), signedTaskResponse)
		if err != nil {
			c.logger.Info("Received error from coordinator", "err", err)
		}

		if response.Success {
			c.logger.Info("Signed task response accepted by coordinator.")
			c.metrics.IncNumTasksAcceptedByCoordinator()
			return
		}
		c.logger.Infof("Retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
	c.logger.Error("Could not send signed task response to coordinator. Tried 5 times.")
}
