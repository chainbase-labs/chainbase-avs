package node

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"

	"github.com/chainbase-labs/chainbase-avs/node/bindings"
)

func (n *ManuscriptNode) StakeIntoStaking(amount *big.Int) error {
	cTokenAddress := common.HexToAddress(n.config.CContractAddress)
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)
	stakeAmount := new(big.Int).Mul(amount, big.NewInt(params.Ether))

	CToken, err := bindings.NewChainbaseToken(cTokenAddress, n.chainbaseClient)
	if err != nil {
		return errors.Wrap(err, "failed to create chainbase token binding")
	}

	txOpts, err := n.chainbaseTxMgr.GetNoSendTxOpts()
	tx, err := CToken.Approve(txOpts, stakingAddress, stakeAmount)
	if err != nil {
		return errors.Wrap(err, "Failed to approve c token transfer")
	}

	_, err = n.chainbaseTxMgr.Send(context.Background(), tx, true)
	if err != nil {
		return errors.Wrap(err, "Failed to send approve tx")
	}

	staking, err := bindings.NewStaking(stakingAddress, n.chainbaseClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err = staking.Stake(txOpts, stakeAmount)
	if err != nil {
		return errors.Wrap(err, "failed to create stake tx")
	}

	receipt, err := n.chainbaseTxMgr.Send(context.Background(), tx, true)
	if err != nil {
		return errors.Wrap(err, "failed to send stake tx")
	}

	n.logger.Infof("stake %s C token into staking contract tx hash %s", amount.String(), receipt.TxHash.String())

	return nil
}

func (n *ManuscriptNode) UnstakeFromStaking() error {
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)

	txOpts, err := n.chainbaseTxMgr.GetNoSendTxOpts()
	staking, err := bindings.NewStaking(stakingAddress, n.chainbaseClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err := staking.Unstake(txOpts)
	if err != nil {
		return errors.Wrap(err, "failed to create unstake tx")
	}

	receipt, err := n.chainbaseTxMgr.Send(context.Background(), tx, true)
	if err != nil {
		return errors.Wrap(err, "failed to send unstake tx")
	}

	n.logger.Infof("unstake from staking contract tx hash %s", receipt.TxHash.String())

	return nil
}

func (n *ManuscriptNode) WithdrawFromStaking() error {
	stakingAddress := common.HexToAddress(n.config.StakingContractAddress)

	txOpts, err := n.chainbaseTxMgr.GetNoSendTxOpts()
	staking, err := bindings.NewStaking(stakingAddress, n.chainbaseClient)
	if err != nil {
		return errors.Wrap(err, "failed to create staking binding")
	}

	tx, err := staking.WithdrawStake(txOpts)
	if err != nil {
		return errors.Wrap(err, "failed to create withdraw tx")
	}

	receipt, err := n.chainbaseTxMgr.Send(context.Background(), tx, true)
	if err != nil {
		return errors.Wrap(err, "failed to send withdraw tx")
	}

	n.logger.Infof("withdraw from staking contract tx hash %s", receipt.TxHash.String())

	return nil
}
