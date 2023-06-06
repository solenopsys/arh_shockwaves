package build

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "build [command]",
	Short: "Build modules (frontend ,container, helm,...)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdHelm)
	Cmd.AddCommand(cmdContainer)
	Cmd.AddCommand(cmdFrontlib)
	Cmd.AddCommand(cmdMicroFrontend)
	Cmd.AddCommand(cmdFrontend)
}
