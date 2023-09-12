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
		var commands []jobs_env.AppCheck
		commands = append(commands, *jobs_env.NewAppCheck("git", []string{"git", "version"}))
		commands = append(commands, *jobs_env.NewAppCheck("pnpm", []string{"pnpm", "-v"}))
		commands = append(commands, *jobs_env.NewAppCheck("go", []string{"go", "version"}))
		commands = append(commands, *jobs_env.NewAppCheck("ng-packagr", []string{"ng-packagr", "-v"}))
		commands = append(commands, *jobs_env.NewAppCheck("nerdctl", []string{"nerdctl", "version"}))

		var ex []jobs.Job

		for _, printableJob := range commands {
			var job jobs.Job = &printableJob
			ex = append(ex, job)
		}
		jobs.ExecuteJobsSync(ex)
	},
}
