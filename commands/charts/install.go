package charts

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdInstall = &cobra.Command{
	Use:   "install [chart] [version] [repository]",
	Short: "Install chart",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ConnectToKubernets()
	},
}
