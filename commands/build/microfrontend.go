package build

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/services"
)

var cmdMicroFrontend = &cobra.Command{
	Use:   "microfrontend [name]",
	Short: "Microfrontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		groupDir := "modules"

		mod, extractError := services.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			println("Error", extractError.Error())
			return
		}

		arg := "run " + mod.Directory + ":build:production"
		argsSplit := strings.Split(arg, " ")

		stdPrinter := services.StdPrinter{Out: make(chan string), Command: "nx", Args: argsSplit}
		go stdPrinter.Processing()
		stdPrinter.Start()
	},
}