package code

import (
	"github.com/spf13/cobra"
)

var cmdSync = &cobra.Command{
	Use:   "sync ",
	Short: "Sync modules by configuration",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//sectionName := args[0]
		//
		//manager, err := funcs.NewWsManager()
		//if err != nil {
		//	io.Fatal(err)
		//}
		//state := manager.GetSectionState(sectionName)
		//if state == "enabled" {
		//	pt := tools.NewPathTools()
		//
		//	//todo refactor
		//	if sectionName == "frontends" {
		//		pt.MoveTo(sectionName)
		//		configs.NewFrontLoader(".", true).SyncFunc()
		//	} else if sectionName == "backends" {
		//		pt.MoveTo(sectionName)
		//		configs.NewBackLoader(".").SyncFunc()
		//	} else {
		//		var name = "xs-treerepo.json" // todo move to const
		//		io.Println("Invalid " + name + ", config type only xs-fronts or xs-backs allowed")
		//		return
		//	}
		//
		//	pt.MoveToBasePath()
		//} else {
		//	io.Println("Invalid argument")
		//	return
		//}

	},
}
