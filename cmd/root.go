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
	rootCmd.AddCommand(registerCmd, runCmd)

	BindRegisterConfig(registerCmd, &cfg)
	BindRegisterConfig(runCmd, &cfg)
}

func Execute() error {
	return rootCmd.Execute()
}
