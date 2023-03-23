package user

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "user [command]",
	Short: "User authorisation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cmdLogin)
	Cmd.AddCommand(cmdLogout)
	Cmd.AddCommand(cmdStatus)
}
