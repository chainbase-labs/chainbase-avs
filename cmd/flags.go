package cmd

import (
	"github.com/spf13/cobra"
)

type RegConfig struct {
	AVSAddr    string
	ConfigFile string
}

func BindRegisterConfig(cmd *cobra.Command, cfg *RegConfig) {
	BindAVSAddress(cmd, &cfg.AVSAddr)

	const flagConfig = "config-file"
	cmd.Flags().StringVar(&cfg.ConfigFile, flagConfig, cfg.ConfigFile, "Path to the Eigen-Layer operator yaml configuration file")
	_ = cmd.MarkFlagRequired(flagConfig)
}

func BindAVSAddress(cmd *cobra.Command, addr *string) {
	const flagConfig = "avs-address"
	cmd.Flags().StringVar(addr, flagConfig, *addr, "Optional address of the Omni AVS contract")
	// _ = cmd.MarkFlagRequired(flagConfig)
}
