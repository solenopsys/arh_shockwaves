package node

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdNodeRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		isServer := true
		const command = "sh"
		if isServer {
			utils.CommandApplyFromFile("/usr/local/bin/k3s-uninstall.sh", command)
		} else {
			utils.CommandApplyFromFile("/usr/local/bin/k3s-agent-uninstall.sh", command)
		}
	},
}
