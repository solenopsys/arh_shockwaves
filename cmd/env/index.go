package env

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var Cmd = &cobra.Command{
	Use:   "env [command]",
	Short: "Environment programms",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdInstall)
	Cmd.AddCommand(cmdStatus)
}
