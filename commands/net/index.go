package net

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "net [command]",
	Short: "Solenopsys network information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdList)

}
