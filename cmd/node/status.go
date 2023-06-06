package node

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var cmdNodeStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("NOT IMPLEMENTED")
	},
}
