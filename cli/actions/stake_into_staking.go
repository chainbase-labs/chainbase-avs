package actions

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/core/config"
	"github.com/chainbase-labs/chainbase-avs/node"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

func StakeIntoStaking(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := utils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}

	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	manuscriptNode, err := node.NewNodeFromConfig(nodeConfig, true)
	if err != nil {
		return err
	}

	amountStr := ctx.String("amount")
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		log.Println("Failed converting amount to big.Int")
		return err
	}

	err = manuscriptNode.StakeIntoStaking(amount)
	if err != nil {
		log.Fatalf("Failed to stake c token into staking contract %v", err.Error())
		return err
	}

	return nil
}
