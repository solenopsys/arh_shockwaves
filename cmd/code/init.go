package code

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Workspace initialization",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		jobs.ExecuteOneSync(jobs_fetch.NewTemplateLoad("@solenopsys/tp-workspace", "."))
	},
}
