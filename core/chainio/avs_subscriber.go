package chainio

import (
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
	"github.com/chainbase-labs/chainbase-avs/core/config"
)

type IAvsSubscriber interface {
	SubscribeToNewTasks(newTaskCreatedChan chan *bindings.ChainbaseServiceManagerNewTaskCreated) event.Subscription
	SubscribeToTaskResponses(taskResponseLogs chan *bindings.ChainbaseServiceManagerTaskResponded) event.Subscription
	ParseTaskResponded(rawLog types.Log) (*bindings.ChainbaseServiceManagerTaskResponded, error)
}

// AvsSubscriber Subscribers use a ws connection instead of http connection like Readers
type AvsSubscriber struct {
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return BuildAvsSubscriber(
		config.RegistryCoordinatorAddr,
		config.OperatorStateRetrieverAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethClient *ethclient.Client, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}
}

func (s *AvsSubscriber) SubscribeToNewTasks(newTaskCreatedChan chan *bindings.ChainbaseServiceManagerNewTaskCreated) event.Subscription {
	sub, err := s.AvsContractBindings.ServiceManager.WatchNewTaskCreated(
		&bind.WatchOpts{}, newTaskCreatedChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new ServiceManager tasks", "err", err)
	}
	s.logger.Infof("Subscribed to new ServiceManager tasks")
	return sub
}

func (s *AvsSubscriber) SubscribeToTaskResponses(taskResponseChan chan *bindings.ChainbaseServiceManagerTaskResponded) event.Subscription {
	sub, err := s.AvsContractBindings.ServiceManager.WatchTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to TaskResponded events", "err", err)
	}
	s.logger.Infof("Subscribed to TaskResponded events")
	return sub
}

func (s *AvsSubscriber) ParseTaskResponded(rawLog types.Log) (*bindings.ChainbaseServiceManagerTaskResponded, error) {
	return s.AvsContractBindings.ServiceManager.ChainbaseServiceManagerFilterer.ParseTaskResponded(rawLog)
}
