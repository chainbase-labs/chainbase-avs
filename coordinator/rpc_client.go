package coordinator

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	nodepb "github.com/chainbase-labs/chainbase-avs/api/grpc/node"
	"github.com/chainbase-labs/chainbase-avs/metrics"
)

type ManuscriptRpcClienter interface {
	CreateNewTask(task *nodepb.NewTaskRequest)
}

type ManuscriptRpcClient struct {
	rpcClient            nodepb.ManuscriptNodeServiceClient
	metrics              metrics.Metrics
	logger               logging.Logger
	nodeServerIpPortAddr string
}

func NewManuscriptRpcClient(nodeServerIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*ManuscriptRpcClient, error) {
	return &ManuscriptRpcClient{
		rpcClient:            nil,
		metrics:              metrics,
		logger:               logger,
		nodeServerIpPortAddr: nodeServerIpPortAddr,
	}, nil
}

func (c *ManuscriptRpcClient) dialManuscriptRpcClient() error {
	client, err := grpc.NewClient(c.nodeServerIpPortAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.rpcClient = nodepb.NewManuscriptNodeServiceClient(client)
	return nil
}

// CreateNewTask sends a new task to the Manuscript node.
func (c *ManuscriptRpcClient) CreateNewTask(task *nodepb.NewTaskRequest) {
	if c.rpcClient == nil {
		c.logger.Info("rpc client is nil. Dialing manuscript node rpc client")
		err := c.dialManuscriptRpcClient()
		if err != nil {
			c.logger.Error("Could not dial manuscript rpc client. Cannot send new task to Manuscript node.", "err", err)
			return
		}
	}

	c.logger.Info("Sending new task to manuscript node", "node", c.nodeServerIpPortAddr)
	for i := 0; i < 3; i++ {
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		//defer cancel()

		response, err := c.rpcClient.ReceiveNewTask(context.Background(), task)
		if err != nil {
			c.logger.Info("Received error from manuscript node", "node", c.nodeServerIpPortAddr, "err", err)
			continue
		}

		if response.Success {
			c.logger.Info("New task accepted by manuscript node.", "node", c.nodeServerIpPortAddr)
			return
		}
		c.logger.Info("Retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
	c.logger.Info("Could not send new task to manuscript node. Tried 3 times.")
}
