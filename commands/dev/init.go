package dev

import (
	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init ",
	Short: "Init tree repository",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loader := NewLoader()
		loader.loadBase()
		NewHelper().initRepository()
		loader.syncModules()
	},
}
