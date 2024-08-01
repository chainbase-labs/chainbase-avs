package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// 导入生成的合约包
)

var (
	cfgFile string // toml config file
)

var rootCmd = &cobra.Command{
	Use:   "chianbase-node",
	Short: "chianbase-node",
	Long:  `chianbase-node`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	Register(cmd.Context(), cfg)
	// },
}

func init() {
	rootCmd.AddCommand(registerCmd, runCmd)

	//toml config file
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "avs.toml", "config file")
}

func Execute() error {
	return rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName("." + "avs")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

}
