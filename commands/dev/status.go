package dev

import (
	"github.com/spf13/cobra"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Show status of installed env programs (git,nx,npm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		println("NOT IMPLEMENTED")
	},
}
