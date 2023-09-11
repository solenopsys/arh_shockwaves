package build

import (
	"github.com/spf13/cobra"
	"xs/internal/extractors"
	jobs_build "xs/internal/jobs/jobs-build"
	"xs/internal/services"
	"xs/pkg/io"
)

var Cmd = &cobra.Command{
	Use:   "build [module/pattern]",
	Short: "Build modules (frontend, module ,container, helm,...)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filter := args[0]

		const PUBLISH = "publish"

		publish := len(args) > 1 && args[1] == PUBLISH

		base := map[string]string{}
		if publish {
			base["publish"] = "true"
		}

		contr := services.UniversalCompileController{
			Executor:  jobs_build.Microfronted{PrintConsole: false},
			Extractor: extractors.Microfrontend{},
		}
		err := contr.SelectLibs(filter)
		if err == nil {
			contr.CompileSelectedLibs()
		} else {
			io.Panic(err)
		}
	},
}
