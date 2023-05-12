package build

import "github.com/spf13/cobra"

var cmdContainer = &cobra.Command{
	Use:   "container [name]",
	Short: "Containers for module build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		chart := args[0]
		println("NOT IMPLEMENTED", chart)
	},
}
