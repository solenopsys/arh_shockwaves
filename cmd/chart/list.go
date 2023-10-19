package chart

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/ui"
	"xs/pkg/wrappers"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List chart",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		kubernetes := wrappers.Kuber{}

		config, err := kubernetes.GetConfig()
		if err != nil {
			io.Panic(err)
		}
		api := wrappers.NewAPI(config)
		charts, err := api.ListHelmCharts("")

		if err != nil {
			io.Panic(err)
		}

		resList := []jobs.ItemTitle{}

		for _, item := range charts.Items {
			resList = append(resList, jobs.ItemTitle{
				Style:       jobs.DEFAULT_STYLE,
				Description: item.Name,
				Name:        item.Name,
				Key:         "chart-list-" + item.Name,
			})
		}

		ui.ListView(resList, "Charts")
	},
}
