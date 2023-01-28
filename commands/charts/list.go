package charts

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdList = cobra.Command{
	Use:   "list",
	Short: "List charts",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ConnectToKubernets()
	},
}
