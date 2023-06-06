package ws

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/internal/funcs"
	"xs/pkg/tools"
)

var cmdLoad = &cobra.Command{
	Use:   "load",
	Short: "Load section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sectionName := args[0]

		manager := funcs.NewWsManager()

		state := manager.GetSectionState(sectionName)
		if state == "enabled" {
			println(sectionName + " already loaded")
			return
		} else if state == "disabled" {
			repository := manager.GetSectionRepository(sectionName)
			err := tools.CreateDirs(sectionName)
			if err != nil {
				panic(err.Error())
			}
			pt := tools.NewPathTools()
			pt.MoveTo(sectionName)
			configs.LoadBase(repository)
			pt.MoveToBasePath()
			manager.SetSectionState(sectionName, "enabled")
			manager.Save()
		} else {
			println("Invalid argument")
			return
		}
	},
}
