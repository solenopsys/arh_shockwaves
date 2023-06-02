package ws

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ws [command]",
	Short: "Workspace commands",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdInit)
	Cmd.AddCommand(cmdLoad)
	Cmd.AddCommand(cmdSync)
	Cmd.AddCommand(cmdState)
}
