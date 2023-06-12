package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/internal/extractors"
	"xs/internal/services"
	"xs/pkg/io"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contr := services.UniversalCompileController{
			Executor:  compilers.Helm{PrintConsole: true},
			Extractor: extractors.Backend{},
			GroupDir:  "deployments",
			RepoType:  configs.BACK,
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
