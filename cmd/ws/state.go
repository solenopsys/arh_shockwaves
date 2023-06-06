package ws

import (
	"github.com/spf13/cobra"
	"xs/internal/funcs"
)

var cmdState = &cobra.Command{
	Use:   "state",
	Short: "Workspace sections state",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		manager := funcs.NewWsManager()
		sections := manager.GetSections()
		for name, section := range sections {
			println(name + ": " + section.State)
		}
	},
}
