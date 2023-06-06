package auth

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var cmdLogout = &cobra.Command{
	Use:   "logout",
	Short: "Forget token",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		JWT_SESSIONS.DeleteSessionTempFiles()
		SOLENOPSYS_KEYS.DeleteSessionTempFiles()

		io.Println("Successful logout")
	},
}
