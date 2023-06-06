package chart

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [chart]",
	Short: "Module chart",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kuber := wrappers.Kuber{}

		config, err := kuber.GetConfig()
		if err != nil {
			io.Fatal(err)
		}
		api := wrappers.NewAPI(config)
		api.DeleteHelmChart(args[0])

		io.Println("Removed: ", args[0])
	},
}
