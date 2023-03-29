package auth

import (
	"github.com/spf13/cobra"
	"xs/utils"
)

var Cmd = &cobra.Command{
	Use:   "auth [command]",
	Short: "Authorisation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var JWT_SESSIONS = utils.NewSessionStorage("jwt-sessions-")
var SOLENOPSYS_KEYS = utils.NewSessionStorage("solenopsys-keys-")

func init() {
	Cmd.AddCommand(cmdLogin)
	Cmd.AddCommand(cmdLogout)
	Cmd.AddCommand(cmdStatus)
}
