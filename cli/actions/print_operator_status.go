package actions

import (
	"encoding/json"
	"log"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/core/config"
	"github.com/chainbase-labs/chainbase-avs/node"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

func PrintOperatorStatus(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}

	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	manuscriptNode, err := node.NewNodeFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	err = manuscriptNode.PrintOperatorStatus()
	if err != nil {
		return err
	}

	return nil
}
