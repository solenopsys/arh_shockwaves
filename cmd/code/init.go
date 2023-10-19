package code

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/pkg/ui"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Workspace initialization",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		load := jobs_fetch.NewTemplateLoad("@solenopsys/tp-workspace", ".")
		executor := jobs.NewExecutor([]jobs.PrintableJob{load})
		ui.ProcessingJobs(executor)
	},
}
