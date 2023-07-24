package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/internal/extractors"
	"xs/internal/services"
	"xs/pkg/io"
)

var cmdFrontlib = &cobra.Command{
	Use:   "frontlib [name]",
	Short: "Frontend build",
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
			Executor:  compilers.Frontlib{PrintConsole: false},
			Extractor: extractors.Frontlib{Base: base},
			GroupDir:  "packages",
			RepoType:  configs.FRONT,
		}

		err := contr.SelectLibs(filter)
		if err == nil {
			contr.CompileSelectedLibs()
		} else {
			io.Panic(err)
		}
	},
}
