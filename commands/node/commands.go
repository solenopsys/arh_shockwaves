package node

import (
	"fmt"
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var Cmd = &cobra.Command{
	Use:   "node [command]",
	Short: "Node control functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

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

var cmdNodeStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start removing ")
	},
}

var cmdNodeInstall = &cobra.Command{
	Use:   "install [connect to node]",
	Short: "Install node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start install")
		utils.CommandApplyFromUrl("https://get.k3s.io", "sh")
	},
}

func init() {
	Cmd.AddCommand(cmdNodeRemove)
	Cmd.AddCommand(cmdNodeStatus)
	Cmd.AddCommand(cmdNodeInstall)
}
