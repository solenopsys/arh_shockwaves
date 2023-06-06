package node

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdNodeInstall = &cobra.Command{
	Use:   "install [connect to node]",
	Short: "Install node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("Start install")
		tools.CommandApplyFromUrl("https://get.k3s.io", "sh")
	},
}
