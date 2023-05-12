package build

import (
	"github.com/spf13/cobra"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		module := args[0]
		println("NOT IMPLEMENTED", module)
	},
}
