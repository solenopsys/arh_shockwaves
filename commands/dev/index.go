package dev

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dev [command]",
	Short: "Developer functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

func init() {
	Cmd.AddCommand(cmdInit)
	Cmd.AddCommand(cmdSync)
}
