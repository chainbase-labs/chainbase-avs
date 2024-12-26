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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/chainbase-labs/chainbase-avs/node/bindings"
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

// RegisterOperatorWithAvs Registration specific functions
func (n *ManuscriptNode) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	// hardcode these things for now
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
	socket := n.nodeSocket
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
func (n *ManuscriptNode) UpdateOperatorSocket() error {
	avsRegistryChainWriter := n.avsWriter.AvsRegistryWriter.(*avsregistry.AvsRegistryChainWriter)
	_, err := avsRegistryChainWriter.UpdateSocket(
		context.Background(),
		eigenSdkTypes.Socket(n.nodeSocket),
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

func (n *ManuscriptNode) StakeIntoStaking(amount *big.Int) error {
	cTokenAddress := common.HexToAddress(n.config.CContractAddress)
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)
	stakeAmount := amount.Mul(amount, big.NewInt(10^18))

	chainbaseRpcClient, err := ethclient.Dial(n.config.ChainbaseRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to create chainbase client")
	}

	CToken, err := bindings.NewChainbaseToken(cTokenAddress, chainbaseRpcClient)
	if err != nil {
		return errors.Wrap(err, "failed to create chainbase token binding")
	}

	txOpts, err := n.avsWriter.TxMgr.GetNoSendTxOpts()
	tx, err := CToken.Approve(txOpts, stakingAddress, stakeAmount)
	if err != nil {
		return errors.Wrap(err, "Failed to approve c token transfer")
	}

	_, err = n.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		return errors.Wrap(err, "Failed to send approve tx")
	}

	staking, err := bindings.NewStaking(stakingAddress, chainbaseRpcClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err = staking.Stake(txOpts, stakeAmount)
	if err != nil {
		return errors.Wrap(err, "failed to create stake tx")
	}

	receipt, err := n.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		return errors.Wrap(err, "failed to send stake tx")
	}

	n.logger.Infof("stake %s into staking contract tx hash %s", amount.String(), receipt.TxHash.String())

	return nil
}

func (n *ManuscriptNode) UnstakeFromStaking() error {
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)

	chainbaseRpcClient, err := ethclient.Dial(n.config.ChainbaseRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to create chainbase client")
	}

	txOpts, err := n.avsWriter.TxMgr.GetNoSendTxOpts()
	staking, err := bindings.NewStaking(stakingAddress, chainbaseRpcClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err := staking.Unstake(txOpts)
	if err != nil {
		return errors.Wrap(err, "failed to create unstake tx")
	}

	receipt, err := n.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		return errors.Wrap(err, "failed to send unstake tx")
	}

	n.logger.Infof("unstake from staking contract tx hash %s", receipt.TxHash.String())

	return nil
}

func (n *ManuscriptNode) WithdrawFromStaking() error {
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)

	chainbaseRpcClient, err := ethclient.Dial(n.config.ChainbaseRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to create chainbase client")
	}

	txOpts, err := n.avsWriter.TxMgr.GetNoSendTxOpts()
	staking, err := bindings.NewStaking(stakingAddress, chainbaseRpcClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err := staking.WithdrawStake(txOpts)
	if err != nil {
		return errors.Wrap(err, "failed to create withdraw tx")
	}

	receipt, err := n.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		return errors.Wrap(err, "failed to send withdraw tx")
	}

	n.logger.Infof("withdraw from staking contract tx hash %s", receipt.TxHash.String())

	return nil
}
