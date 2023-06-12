package build

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var Cmd = &cobra.Command{
	Use:   "build [command]",
	Short: "Build modules (frontend ,container, helm,...)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdHelm)
	Cmd.AddCommand(cmdContainer)
	Cmd.AddCommand(cmdFrontlib)
	Cmd.AddCommand(cmdFrontend)
}
