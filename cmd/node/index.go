package node

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var Cmd = &cobra.Command{
	Use:   "node [command]",
	Short: "Node control functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdNodeRemove)
	Cmd.AddCommand(cmdNodeStatus)
	Cmd.AddCommand(cmdNodeInstall)
	Cmd.AddCommand(cmdNodeVars)
}
