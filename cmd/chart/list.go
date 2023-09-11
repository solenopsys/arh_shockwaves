package chart

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_chart "xs/internal/jobs/jobs-chart"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List chart",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		jobs.ExecuteOneSync(jobs_chart.NewChartList())
	},
}
