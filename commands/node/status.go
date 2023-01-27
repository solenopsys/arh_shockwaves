package node

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmdNodeStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start removing ")
	},
}
