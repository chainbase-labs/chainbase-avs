package actions

import (
	"encoding/json"
	"log"
	"os"

	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"

	"github.com/chainbase-labs/chainbase-avs/core/config"
	"github.com/chainbase-labs/chainbase-avs/node"
	"github.com/chainbase-labs/chainbase-avs/node/types"
)

func RegisterOperatorWithAvs(ctx *cli.Context) error {
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

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Printf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}
	operatorEcdsaPrivKey, err := sdkecdsa.ReadKey(
		nodeConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	err = manuscriptNode.RegisterOperatorWithAvs(operatorEcdsaPrivKey)
	if err != nil {
		return err
	}

	return nil
}
