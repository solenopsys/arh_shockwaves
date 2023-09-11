package jobs_chart

import (
	"xs/internal/jobs"
	"xs/pkg/wrappers"
)

type ChartInstall struct {
	chart     string
	repoUrl   string
	version   string
	namespace string
}

func (c *ChartInstall) Execute() *jobs.Result {
	kubernetes := wrappers.Kuber{}
	config, err := kubernetes.GetConfig()
	if err != nil {
		return &jobs.Result{
			Success: false,
			Err:     err,
		}
	}
	api := wrappers.NewAPI(config)

	simple, err := api.CreateHelmChartSimple(c.chart, c.repoUrl, c.version, c.namespace)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Err:     err,
		}
	}

	return &jobs.Result{
		Success:     true,
		Err:         nil,
		Description: "Installed: " + simple.Name,
	}
}

func NewChartInstall(chart string, repoUrl string, version string) *ChartInstall {
	return &ChartInstall{chart: chart, repoUrl: repoUrl, version: version, namespace: "default"}
}
