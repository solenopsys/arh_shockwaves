package build

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/tools"
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

		base := map[string]string{}
		if publish {
			base["publish"] = "true"
		}

		wm, err := configs.GetInstanceWsManager()
		if err != nil {
			io.Panic(err)
		}

		cm := configs.GetInstanceConfManager()

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

		for builderName, libs := range buildGroups {
			jobsPlan := compilers.NewCompilePlanning(publish).GetPlan(builderName, libs)
			io.Println("SECTION:", builderName)
			for _, job := range jobsPlan {
				jobs.PrintJob(job.Description())
			}
		}

		confirm := tools.ConfirmDialog("Build this libraries?")

		if confirm {
			io.Println("Proceeding with the action.")
			for builderName, libs := range buildGroups {
				jobsPlan := compilers.NewCompilePlanning(publish).GetPlan(builderName, libs)
				io.Println("SECTION:", builderName)
				jobs.ExecuteJobsSync(jobs.ConvertJobs(jobsPlan))
			}
		} else {
			io.Println("Canceled.")
		}

		//contr := compilers.CompilerPlanGenerator{
		//	Executor:  jobs_build.Microfronted{PrintConsole: false},
		//	Extractor: extractors.Microfrontend{},
		//}
		//err := contr.SelectLibs(filter)
		//if err == nil {
		//	contr.CompileSelectedLibs()
		//} else {
		//	io.Panic(err)
		//}
	},
}

//*Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
