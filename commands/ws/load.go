package ws

import (
	"github.com/spf13/cobra"
	"xs/services"
	"xs/utils"
)

var cmdLoad = &cobra.Command{
	Use:   "load",
	Short: "Load section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sectionName := args[0]

		manager := utils.NewWsManager()

		state := manager.GetSectionState(sectionName)
		if state == "enabled" {
			println(sectionName + " already loaded")
			return
		} else if state == "disabled" {
			repository := manager.GetSectionRepository(sectionName)
			err := utils.CreateDirs(sectionName)
			if err != nil {
				panic(err.Error())
			}
			pt := utils.NewPathTools()
			pt.MoveTo(sectionName)
			services.LoadBase(repository)
			pt.MoveToBasePath()
			manager.SetSectionState(sectionName, "enabled")
			manager.Save()
		} else {
			println("Invalid argument")
			return
		}
	},
}
