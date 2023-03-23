package user

import (
	"github.com/spf13/cobra"
)

var cmdLogout = &cobra.Command{
	Use:   "logout",
	Short: "Forget token",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
