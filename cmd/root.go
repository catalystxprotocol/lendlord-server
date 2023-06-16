package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "v0.1"
	Date    = "2023-02-16"
)

var rootCmd = &cobra.Command{
	Use:   "LendLord",
	Short: "LendLord NFT Rent Market",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
