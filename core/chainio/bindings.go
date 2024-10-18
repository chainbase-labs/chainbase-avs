package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/chainbase-labs/chainbase-avs/contracts/bindings"
)

type AvsManagersBindings struct {
	ServiceManager *bindings.ChainbaseServiceManager
	ethClient      eth.Client
	logger         logging.Logger
}

func NewAvsManagersBindings(registryCoordinatorAddr, _ gethcommon.Address, ethClient eth.Client, logger logging.Logger) (*AvsManagersBindings, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}

	serviceManagerAddr, err := contractRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	contractServiceManager, err := bindings.NewChainbaseServiceManager(serviceManagerAddr, ethClient)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}

	return &AvsManagersBindings{
		ServiceManager: contractServiceManager,
		ethClient:      ethClient,
		logger:         logger,
	}, nil
}

func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*bindings.ERC20Mock, error) {
	contractErc20Mock, err := bindings.NewERC20Mock(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return contractErc20Mock, nil
}
