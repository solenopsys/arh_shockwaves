package dev

import (
	"github.com/spf13/cobra"
)

var cmdSync = &cobra.Command{
	Use:   "sync ",
	Short: "Sync modules by configuration",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loader := NewLoader()
		loader.syncModules()
	},
}
