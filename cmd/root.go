package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// 导入生成的合约包
)

var (
	cfgFile string // toml config file
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

	//toml config file
	cobra.OnInitialize(initConfig)

}

func Execute() error {
	return rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AutomaticEnv()

}
