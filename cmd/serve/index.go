package serve

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "serve [command]",
	Short: "Serve frontend",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdFront)
}
