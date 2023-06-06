package node

import (
	"github.com/spf13/cobra"
	"xs/pkg/wrappers"
)

var cmdNodeVars = &cobra.Command{
	Use:   "vars",
	Short: "Set variables (namespaces, service accounts, etc.)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		kuber := wrappers.Kuber{}
		err := kuber.CreateNamespace("installers")
		if err != nil {
			panic(err)
		} else {
			println("Namespace installers created")
		}
	},
}
