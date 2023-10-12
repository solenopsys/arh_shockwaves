package env

import (
	"github.com/spf13/cobra"
	"xs/internal/jobs"
	jobs_env "xs/internal/jobs/jobs-env"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Show status of installed env programs (git,pnpm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var jobsPlan = []jobs.Job{
			jobs_env.NewAppCheck("git", []string{"git", "version"}),
			jobs_env.NewAppCheck("pnpm", []string{"pnpm", "-v"}),
			jobs_env.NewAppCheck("go", []string{"go", "version"}),
			jobs_env.NewAppCheck("ng-packagr", []string{"ng-packagr", "-v"}),
			jobs_env.NewAppCheck("nerdctl", []string{"nerdctl", "version"}),
		}
		stat := jobs.ExecuteJobsSync(jobsPlan)

		println("Status:")
		for _, item := range stat {
			if item.Success {
				println("✅ " + item.Description)
			} else {
				println("❌ " + item.Description)
			}
		}
	},
}
