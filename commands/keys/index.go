package keys

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "keys [command]",
	Short: "Keys manipulation functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdPubkey)
	Cmd.AddCommand(cmdSeed)
	Cmd.AddCommand(cmdAccount)
}
