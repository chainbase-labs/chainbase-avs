package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/coordinator"
	"github.com/chainbase-labs/chainbase-avs/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {
	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "chainbase-coordinator"
	app.Usage = "Chainbase Coordinator"
	app.Description = "This is a service run by Chainbase. It generates tasks on-chain and aggregates the responses coming from the manuscript node."
	app.Action = coordinatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Coordinator run failed.", "Message:", err)
	}
}

func coordinatorMain(ctx *cli.Context) error {
	log.Println("Initializing Coordinator")
	coordinatorConfig, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(coordinatorConfig, "", "  ")
	if err != nil {
		coordinatorConfig.Logger.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	coordinatorInstance, err := coordinator.NewCoordinator(coordinatorConfig)
	if err != nil {
		return err
	}

	err = coordinatorInstance.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
