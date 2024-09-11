package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/core/config"
	"github.com/chainbase-labs/chainbase-avs/node"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.ConfigFileFlag}
	app.Name = "chainbase-manuscript-node"
	app.Usage = "Chainbase Manuscript Node"
	app.Description = "Service run by operator that receive task from coordinator, calculate, signs, and sends result to the coordinator."

	app.Action = nodeMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func nodeMain(ctx *cli.Context) error {
	log.Println("Initializing manuscript node")
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
	log.Println("initialized manuscript node")

	log.Println("starting manuscript node")
	err = manuscriptNode.Start(context.Background())
	if err != nil {
		return err
	}
	log.Println("started manuscript node")

	return nil

}
