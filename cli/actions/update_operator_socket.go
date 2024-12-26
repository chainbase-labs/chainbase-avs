package actions

import (
	"encoding/json"
	"log"

	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/core"
	"github.com/chainbase-labs/chainbase-avs/core/config"
	"github.com/chainbase-labs/chainbase-avs/node"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

func UpdateOperatorSocket(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := core.ReadYamlConfig(configPath, &nodeConfig)
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

	err = manuscriptNode.UpdateOperatorSocket()
	if err != nil {
		return err
	}

	return nil
}
