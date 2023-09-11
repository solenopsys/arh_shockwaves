package chart

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_chart "xs/internal/jobs/jobs-chart"
)

var cmdInstall = &cobra.Command{
	Use:   "install [chart] [version] [repository]",
	Short: "Install chart",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		chart := args[0]
		repoUrl := args[2]
		version := args[1]
		jobs.ExecuteOneSync(jobs_chart.NewChartInstall(chart, repoUrl, version))
	},
}
