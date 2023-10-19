package code

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/ui"
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

		jobsPlan := jobs.ConvertPjToJi(makeRemovePlan("*"))
		applied, changedPattern := ui.FilteredListView(jobsPlan, "Sources for remove", pattern)
		if applied {
			jobsPlanApplied := makeRemovePlan(changedPattern)
			executor := jobs.NewExecutor(jobsPlanApplied)
			ui.ProcessingJobs(executor)
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
