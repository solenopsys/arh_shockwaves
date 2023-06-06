package chart

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List chart",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		kuber := wrappers.Kuber{}

		config, err := kuber.GetConfig()
		if err != nil {
			io.Fatal(err)
		}
		api := wrappers.NewAPI(config)
		charts, err := api.ListHelmCharts("")

		if err != nil {
			io.Fatal(err)
		}
		for _, item := range charts.Items {
			io.Println(item.Name)
		}

	},
}
