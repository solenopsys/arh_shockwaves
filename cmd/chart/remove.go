package chart

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_chart "xs/internal/jobs/jobs-chart"
	"xs/pkg/ui"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [chart]",
	Short: "Module chart",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		chart := args[0]
		load := jobs_chart.NewChartRemove(chart)

		executor := jobs.NewExecutor([]jobs.PrintableJob{load})
		ui.ProcessingJobs(executor)
	},
}
