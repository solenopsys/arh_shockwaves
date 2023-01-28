package charts

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [chart] [version] [repository]",
	Short: "Module chart",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ConnectToKubernets()
	},
}
