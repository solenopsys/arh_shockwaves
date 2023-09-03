package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/internal/extractors"
	"xs/internal/services"
	"xs/pkg/io"
)

var cmdModule = &cobra.Command{
	Use:   "module [name]",
	Short: "module build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contr := services.UniversalCompileController{
			Executor:  compilers.Microfronted{PrintConsole: false},
			Extractor: extractors.Microfrontend{},
			GroupDir:  "modules",
			RepoType:  configs.FRONT,
		}
		filter := args[0]
		err := contr.SelectLibs(filter)
		if err == nil {
			contr.CompileSelectedLibs()
		} else {
			io.Panic(err)
		}
	},
}
