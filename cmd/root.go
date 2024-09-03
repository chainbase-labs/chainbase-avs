package cmd

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string // toml config file
)

var rootCmd = &cobra.Command{
	Use:   "chainbase-node",
	Short: "chainbase-node",
	Long:  `chainbase-node`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	Register(cmd.Context(), cfg)
	// },
}

func init() {
	rootCmd.AddCommand(registerCmd, runCmd)

	//toml config file
	cobra.OnInitialize(initConfig)

}

func Execute() error {
	return rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if _, err := os.Stat(".env"); err == nil {
		slog.Info("load env from .env file")
		_ = godotenv.Load()
	}
	viper.AutomaticEnv()

}
