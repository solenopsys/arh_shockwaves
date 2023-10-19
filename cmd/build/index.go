package build

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/ui"
)

var Cmd = &cobra.Command{
	Use:   "build [module/pattern] [-d]",
	Short: "Build modules (frontend, module ,container, helm,...)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filter := args[0]

		const PUBLISH = "deploy"

		err := tools.ToWorkspaceRootDir()
		if err != nil {
			io.Fatal("Workspace root dir not found")
		}

		publish := len(args) > 1 && args[1] == PUBLISH
		jobsPlan := jobs.ConvertPjToJi(JobsPlanByGroups("*", publish))
		applied, changedFilter := ui.FilteredListView(jobsPlan, "Build this projects", filter)
		if applied {
			jobsPlanApplied := JobsPlanByGroups(changedFilter, publish)
			executor := jobs.NewExecutor(jobsPlanApplied)
			ui.ProcessingJobs(executor)
		}
	},
}

func JobsPlanByGroups(filter string, publish bool) []jobs.PrintableJob {
	cm := configs.GetInstanceConfManager()
	wm, err := configs.GetInstanceWsManager()
	if err != nil {
		io.Panic(err)
	}

	libs := wm.FilterLibs(filter)

	mapping := cm.GetBuildersMapping()

	buildGroups := make(map[string][]*configs.XsModule)
	for _, lib := range libs {
		for parentDirs, builderName := range mapping {
			if strings.HasPrefix(lib.Directory, parentDirs) {
				buildGroups[builderName] = append(buildGroups[builderName], lib)
			}
		}
	}
	resultPlan := make([]jobs.PrintableJob, 0)
	for builderName, libs := range buildGroups {
		jobsPlan := compilers.NewCompilePlanning(publish).GetPlan(builderName, libs)
		resultPlan = append(resultPlan, jobsPlan...)
	}

	return resultPlan
}
