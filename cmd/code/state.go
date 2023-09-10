package code

import (
	"github.com/spf13/cobra"
)

var cmdState = &cobra.Command{
	Use:   "state",
	Short: "Workspace sections state",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//manager, err := funcs.NewWsManager()
		//if err != nil {
		//	io.Fatal(err)
		//}
		//sections := manager.GetSections()
		//for name, section := range sections {
		//	io.Println(name + ": " + section.State)
		//}
	},
}
