package node

import (
	"github.com/spf13/cobra"
	"xs/pkg/tools"
)

var cmdNodeRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		isServer := true
		const command = "sh"
		if isServer {
			tools.CommandApplyFromFile("/usr/local/bin/k3s-uninstall.sh", command)
		} else {
			tools.CommandApplyFromFile("/usr/local/bin/k3s-agent-uninstall.sh", command)
		}
	},
}
