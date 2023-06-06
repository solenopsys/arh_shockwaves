package ws

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/internal/funcs"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdSync = &cobra.Command{
	Use:   "sync ",
	Short: "Sync modules by configuration",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sectionName := args[0]

		manager := funcs.NewWsManager()
		state := manager.GetSectionState(sectionName)
		if state == "enabled" {
			pt := tools.NewPathTools()

			//todo refactor
			if sectionName == "frontends" {
				pt.MoveTo(sectionName)
				configs.NewFrontLoader("./").SyncFunc()
			} else if sectionName == "backends" {
				pt.MoveTo(sectionName)
				configs.NewBackLoader("./").SyncFunc()
			} else {
				io.Println("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
				return
			}

			pt.MoveToBasePath()
		} else {
			io.Println("Invalid argument")
			return
		}

	},
}
