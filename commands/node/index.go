package node

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "node [command]",
	Short: "Node control functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdNodeRemove)
	Cmd.AddCommand(cmdNodeStatus)
	Cmd.AddCommand(cmdNodeInstall)
}
