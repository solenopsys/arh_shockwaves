package node

import (
	"fmt"
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdNodeInstall = &cobra.Command{
	Use:   "install [connect to node]",
	Short: "Install node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start install")
		utils.CommandApplyFromUrl("https://get.k3s.io", "sh")
	},
}
