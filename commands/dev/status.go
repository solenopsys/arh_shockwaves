package dev

import (
	"github.com/spf13/cobra"
	"xs/services"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Show status of installed env programs (git,nx,npm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loader := services.NewLoader()
		loader.SyncModules()
	},
}
