package ws

import (
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/os"
	"xs/services"
	"xs/utils"
)

var cmdSync = &cobra.Command{
	Use:   "sync ",
	Short: "Sync modules by configuration",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "./xs.json"
		exists := os.FileExists(fileName)
		if exists {
			config := services.LoadConfigFile(fileName)
			repoType := services.FileTypeMapping[config.Format.Name]

			sectionName := args[0]

			manager := utils.NewWsManager()

			state := manager.GetSectionState(sectionName)
			if state == "enabled" {
				if repoType == "frontends" {
					services.NewFrontLoader().SyncFunc()
				} else if repoType == "backends" {
					services.NewBackLoader().SyncFunc()
				} else {
					println("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
					return
				}
			} else {
				println("Invalid argument")
				return
			}

		} else {
			println("xs.json not found, directory not initialized")
		}
	},
}
