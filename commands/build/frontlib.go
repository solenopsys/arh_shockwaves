package build

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"xs/services"
)

var cmdFrontlib = &cobra.Command{
	Use:   "frontlib [name]",
	Short: "Frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		groupDir := "libraries"

		mod, extractError := services.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			println("Error", extractError.Error())
			return
		}

		arg := "build"
		argsSplit := strings.Split(arg, " ")

		path := "./" + groupDir + "/" + mod.Directory

		errDir := os.Chdir(path)
		if errDir != nil {
			fmt.Println(errDir)
			return
		}

		stdPrinter := services.StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit}
		go stdPrinter.Processing()
		stdPrinter.Start()
	},
}
