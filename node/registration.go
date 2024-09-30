package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"

	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (n *ManuscriptNode) RegisterOperatorWithEigenlayer() error {
	op := eigenSdkTypes.Operator{
		Address:                 n.operatorAddr.String(),
		EarningsReceiverAddress: n.operatorAddr.String(),
	}
	_, err := n.eigenlayerWriter.RegisterAsOperator(context.Background(), op)
	if err != nil {
		n.logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}
	n.logger.Infof("Registered operator with eigenlayer")

	return nil
}

func (n *ManuscriptNode) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
	_, tokenAddr, err := n.eigenlayerReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
	if err != nil {
		n.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	contractErc20Mock, err := n.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		n.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := n.avsWriter.TxMgr.GetNoSendTxOpts()
	tx, err := contractErc20Mock.Mint(txOpts, n.operatorAddr, amount)
	if err != nil {
		n.logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = n.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		n.logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = n.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
	if err != nil {
		n.logger.Error("Error depositing into strategy", "err", err)
		return err
	}
	n.logger.Infof("Deposited %s into strategy %s", amount, strategyAddr)

	return nil
}

// RegisterOperatorWithAvs Registration specific functions
func (n *ManuscriptNode) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	// hardcode these things for now
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
	socket := n.nodeServerIpPortAddr
	_, err := n.avsWriter.RegisterOperator(
		context.Background(),
		operatorEcdsaKeyPair,
		n.blsKeypair,
		quorumNumbers,
		socket,
	)
	if err != nil {
		n.logger.Errorf("Unable to register operator with avs registry coordinator")
		return err
	}
	n.logger.Infof("Registered operator with avs registry coordinator")

	return nil
}

// OperatorStatus print status of operator
type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func (n *ManuscriptNode) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	operatorId, err := n.avsReader.GetOperatorId(&bind.CallOpts{}, n.operatorAddr)
	if err != nil {
		return err
	}
	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := n.operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      n.operatorAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          n.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          n.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(n.operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}

// UpdateOperatorSocket update operator socket
func (n *ManuscriptNode) UpdateOperatorSocket(socket string) error {
	avsRegistryChainWriter := n.avsWriter.AvsRegistryWriter.(*avsregistry.AvsRegistryChainWriter)
	_, err := avsRegistryChainWriter.UpdateSocket(
		context.Background(),
		eigenSdkTypes.Socket(socket),
	)
	if err != nil {
		n.logger.Errorf("Unable to update operator socket")
		return err
	}
	n.logger.Infof("Update operator socket successfully")

	return nil
}

// DeregisterOperatorWithAvs deregister operator with avs
func (n *ManuscriptNode) DeregisterOperatorWithAvs() error {
	// hardcode these things for now
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
	pubkey := utils.ConvertToBN254G1Point(n.blsKeypair.GetPubKeyG1())
	_, err := n.avsWriter.DeregisterOperator(
		context.Background(),
		quorumNumbers,
		pubkey,
	)
	if err != nil {
		n.logger.Errorf("Unable to deregister operator with avs registry coordinator")
		return err
	}
	n.logger.Infof("Deregistered operator with avs registry coordinator")

	return nil
}
