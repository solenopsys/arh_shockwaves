package cluster

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Cluster status",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ConnectToKubernets()
	},
}
