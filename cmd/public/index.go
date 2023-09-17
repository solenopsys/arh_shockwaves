package public

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "public [command]",
	Short: "Public content in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdFile)
	Cmd.AddCommand(cmdDir)
	Cmd.AddCommand(cmdGit)
	Cmd.AddCommand(cmdSyncGit)
	Cmd.AddCommand(cmdName)
}
