package auth

import (
	"github.com/spf13/cobra"
	"xs/internal/storage"
)

var Cmd = &cobra.Command{
	Use:   "auth [command]",
	Short: "Authorisation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var JWT_SESSIONS = storage.NewSessionStorage("jwt-sessions-")
var SOLENOPSYS_KEYS = storage.NewSessionStorage("solenopsys-keys-")

func init() {
	Cmd.AddCommand(cmdLogin)
	Cmd.AddCommand(cmdLogout)
	Cmd.AddCommand(cmdStatus)
}
