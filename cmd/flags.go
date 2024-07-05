package cmd

import (
	"github.com/spf13/cobra"
)

type RegConfig struct {
	AVSAddr    string
	PrivateKey string
}

func BindRegConfig(cmd *cobra.Command, cfg *RegConfig) {
	BindAVSAddress(cmd, &cfg.AVSAddr)

	const flagConfig = "private-key"
	cmd.Flags().StringVar(&cfg.PrivateKey, flagConfig, cfg.PrivateKey, "private key of the operator")
	// _ = cmd.MarkFlagRequired(flagConfig)
}

func BindAVSAddress(cmd *cobra.Command, addr *string) {
	const flagConfig = "avs-address"
	cmd.Flags().StringVar(addr, flagConfig, *addr, "Optional address of the Omni AVS contract")
	// _ = cmd.MarkFlagRequired(flagConfig)
}
