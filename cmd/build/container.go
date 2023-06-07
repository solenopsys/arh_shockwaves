package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/pkg/io"
)

var cmdContainer = &cobra.Command{
	Use:   "container [name]",
	Short: "Containers for module build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]

		hce := compilers.ContainerCompileExecutor{PrintConsole: true}

		err := hce.Compile(map[string]string{
			"name": name,
		})

		if err != nil {
			io.Println("Error", err.Error())
		}
	},
}
