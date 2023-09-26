package code

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdState = &cobra.Command{
	Use:   "remove [pattern]",
	Short: "Workspace sections state",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]

		err := tools.ToWorkspaceRootDir()
		if err != nil {
			io.Fatal("Workspace root dir not found")
		}

		jobsPlan := makeRemovePlan(pattern)

		for _, job := range jobsPlan {
			jobs.PrintJob(job.Description())
		}

		confirm := tools.ConfirmDialog("Remove packets?")

		if confirm {
			io.Println("Proceeding with the action.")
			jobs.ExecuteJobsSync(jobs.ConvertJobs(jobsPlan))
		} else {
			io.Println("Canceled.")
		}
	},
}

func makeRemovePlan(pattern string) []jobs.PrintableJob {
	processorsManager := CreateProcessors([]string{"code", "remove"})
	codeJobs := make([]jobs.PrintableJob, 0)
	confManager, err := configs.GetInstanceWsManager()
	if err != nil {
		io.Fatal("Workspace root dir not found")
	}

	libs := confManager.FilterLibs(pattern)

	for _, lib := range libs {
		subDir := strings.Split(lib.Directory, "/")[0]
		preJobs := processorsManager.GetPreProcessors(subDir, lib.Name, lib.Directory)
		postJobs := processorsManager.GetPostProcessors(subDir, lib.Name, lib.Directory)
		codeJobs = append(codeJobs, preJobs...)
		codeJobs = append(codeJobs, jobs_fetch.NewCodeRemove(lib.Name, lib.Directory))
		codeJobs = append(codeJobs, postJobs...)
	}

	return codeJobs
}
