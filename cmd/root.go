package cmd

import (
	"github.com/spf13/cobra"
	// 导入生成的合约包
)

var (
	cfg RegConfig
)

var rootCmd = &cobra.Command{
	Use:   "chianbase-avs",
	Short: "chianbase-avs",
	Long:  `chianbase-avs`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	Register(cmd.Context(), cfg)
	// },
}

func init() {
	BindRegisterConfig(rootCmd, &cfg)
}

func Execute() error {
	return rootCmd.Execute()
}

// Register registers the operator with the ChainBase AVS contract.
//
// It assumes that the operator is already registered with the Eigen-Layer
// and that the eigen-layer configuration file (and ecdsa keystore) is present on disk.
