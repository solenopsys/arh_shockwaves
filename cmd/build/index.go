package build

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var Cmd = &cobra.Command{
	Use:   "build [module/pattern]",
	Short: "Build modules (frontend, module ,container, helm,...)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filter := args[0]

		const PUBLISH = "publish"

		err := tools.ToWorkspaceRootDir()
		if err != nil {
			io.Fatal("Workspace root dir not found")
		}

		publish := len(args) > 1 && args[1] == PUBLISH

		base := map[string]string{}
		if publish {
			base["publish"] = "true"
		}

		wm, err := configs.NewWsManager()
		if err != nil {
			io.Panic(err)
		}

		cm, err := configs.NewConfigurationManager()

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

		jobs := compilers.NewCompilePlanning().GetPlan("frontlib", buildGroups["frontlib"])

		for _, job := range jobs {
			println(job.Description())
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
