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

func (c ChartInstall) Execute() *jobs.Result {
	kubernetes := wrappers.Kuber{}
	config, err := kubernetes.GetConfig()
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	api := wrappers.NewAPI(config)

	simple, err := api.CreateHelmChartSimple(c.chart, c.repoUrl, c.version, c.namespace)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Installed: " + simple.Name,
	}
}
func (c ChartInstall) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: c.repoUrl,
		Name:        c.chart,
		Key:         "chart-install-" + c.chart,
	}
}

func NewChartInstall(chart string, repoUrl string, version string) ChartInstall {
	return ChartInstall{chart: chart, repoUrl: repoUrl, version: version, namespace: "default"}
}
