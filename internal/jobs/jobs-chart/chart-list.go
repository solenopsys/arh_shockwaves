package jobs_chart

import (
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type ChartList struct {
}

func (c *ChartList) Execute() *jobs.Result {

	kubernetes := wrappers.Kuber{}

	config, err := kubernetes.GetConfig()
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	api := wrappers.NewAPI(config)
	charts, err := api.ListHelmCharts("")

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	for _, item := range charts.Items {
		io.Println(item.Name)
	}

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Charts count: " + string(len(charts.Items)),
	}
}

func NewChartList() *ChartList {
	return &ChartList{}
}
