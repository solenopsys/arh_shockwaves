package build

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
	io "xs/pkg/io"
)

var cmdFrontend = &cobra.Command{
	Use:   "frontend [name]",
	Short: "frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		groupDir := "modules"

		mod, extractError := configs.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			io.Println("Error", extractError.Error())
			return
		}

		arg := "ng " + mod.Directory + ":build:production"
		argsSplit := strings.Split(arg, " ")

		stdPrinter := io.StdPrinter{Out: make(chan string), Command: "nx", Args: argsSplit}
		go stdPrinter.Processing()
		stdPrinter.Start()
	},
}
