package chart

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "chart [command]",
	Short: "Charts manipulation functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdInstall)
	Cmd.AddCommand(cmdRemove)
	Cmd.AddCommand(cmdList)

}
