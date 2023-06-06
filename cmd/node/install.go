package node

import (
	"fmt"
	"github.com/spf13/cobra"
	"xs/pkg/tools"
)

var cmdNodeInstall = &cobra.Command{
	Use:   "install [connect to node]",
	Short: "Install node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start install")
		tools.CommandApplyFromUrl("https://get.k3s.io", "sh")
	},
}
