package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	io "xs/pkg/io"
)

var cmdFrontend = &cobra.Command{
	Use:   "frontend [name]",
	Short: "frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		hce := compilers.AngularFrontendCompileExecutor{PrintConsole: true}

		err := hce.Compile(map[string]string{
			"name": name,
		})

		if err != nil {
			io.Println("Error", err.Error())
		}
	},
}
