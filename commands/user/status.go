package user

import (
	"github.com/spf13/cobra"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of user",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
