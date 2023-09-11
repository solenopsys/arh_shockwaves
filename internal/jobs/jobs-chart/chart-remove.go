package jobs_chart

import (
	"xs/internal/jobs"
	"xs/pkg/wrappers"
)

type ChartRemove struct {
	chart     string
	namespace string
}

func (c *ChartRemove) Execute() *jobs.Result {

	kubernetes := wrappers.Kuber{}

	config, err := kubernetes.GetConfig()
	if err != nil {
		return &jobs.Result{
			Success: false,
			Err:     err,
		}
	}
	api := wrappers.NewAPI(config)
	err = api.DeleteHelmChart(c.chart)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Err:     err,
		}
	}

	return &jobs.Result{
		Success:     true,
		Err:         nil,
		Description: "Removed: " + c.chart,
	}
}

func NewChartRemove(chart string) *ChartInstall {
	return &ChartInstall{chart: chart, namespace: "default"}
}
