package ws

import (
	"github.com/spf13/cobra"
	"xs/services"
	"xs/utils"
)

var cmdSync = &cobra.Command{
	Use:   "sync ",
	Short: "Sync modules by configuration",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sectionName := args[0]

		manager := utils.NewWsManager()
		state := manager.GetSectionState(sectionName)
		if state == "enabled" {
			pt := utils.NewPathTools()

			//todo refactor
			if sectionName == "frontends" {
				pt.MoveTo(sectionName)
				services.NewFrontLoader("./").SyncFunc()
			} else if sectionName == "backends" {
				pt.MoveTo(sectionName)
				services.NewBackLoader("./").SyncFunc()
			} else {
				println("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
				return
			}

			pt.MoveToBasePath()
		} else {
			println("Invalid argument")
			return
		}

	},
}
