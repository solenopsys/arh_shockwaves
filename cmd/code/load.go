package code

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/internal/funcs"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdLoad = &cobra.Command{
	Use:   "load",
	Short: "Tags section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sectionName := args[0]

		manager, err := funcs.NewWsManager()
		if err != nil {
			io.Fatal(err)
		}
		state := manager.GetSectionState(sectionName)
		if state == "enabled" {
			io.Println(sectionName + " already loaded")
			return
		} else if state == "disabled" {
			repository := manager.GetSectionRepository(sectionName)
			err := tools.CreateDirs(sectionName)
			if err != nil {
				io.Panic(err)
			}
			pt := tools.NewPathTools()
			pt.MoveTo(sectionName)
			configs.LoadBase(repository)
			pt.MoveToBasePath()
			manager.SetSectionState(sectionName, "enabled")
			manager.Save()
		} else {
			io.Println("Invalid argument")
			return
		}
	},
}
