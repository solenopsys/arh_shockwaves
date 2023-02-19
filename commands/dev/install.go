package dev

import (
	"github.com/spf13/cobra"
	"xs/services"
)

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install all necessary programs (git,nx,npm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loader := services.NewLoader()
		loader.SyncModules()
	},
}
