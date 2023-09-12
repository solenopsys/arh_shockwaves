package code

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "code [command]",
	Short: "Workspace commands",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdInit)
	Cmd.AddCommand(cmdAdd)
	Cmd.AddCommand(cmdSync)
	Cmd.AddCommand(cmdState)
}
