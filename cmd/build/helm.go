package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/pkg/io"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		hce := compilers.HelmCompileExecutor{PrintConsole: true}

		err := hce.Compile(map[string]string{
			"name": name,
		})

		if err != nil {
			io.Println("Error", err.Error())
		}
	},
}
