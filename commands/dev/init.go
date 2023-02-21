package dev

import (
	"github.com/spf13/cobra"
	"xs/services"
)

var cmdInit = &cobra.Command{
	Use:   "init [front/back]",
	Short: "Init monorepo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//validate arg 0 is front or back
		repoType := args[0]

		if repoType == "front" {
			services.LoadBase("https://github.com/solenopsys/xs-fronts-template.git")
			services.NewFrontLoader().SyncFunc()
		} else if repoType == "back" {
			services.LoadBase("https://github.com/solenopsys/xs-backs-template.git")
			services.NewBackLoader().SyncFunc()
		} else {
			println("Invalid argument, only front or back allowed")
			return
		}

	},
}
