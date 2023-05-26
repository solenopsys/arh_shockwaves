package build

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
	"xs/services"
)

func processLib(m string) error {
	groupDir := "libraries"

	mod, extractError := services.ExtractModule(m, groupDir, "front")
	if extractError != nil {
		return extractError
	}
	arg := "build"
	argsSplit := strings.Split(arg, " ")

	path := "./" + groupDir + "/" + mod.Directory

	errDir := os.Chdir(path)
	if errDir != nil {
		return extractError
	}

	stdPrinter := services.StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit}
	go stdPrinter.Processing()
	stdPrinter.Start()
	return nil
}

var cmdFrontlib = &cobra.Command{
	Use:   "frontlib [name]",
	Short: "Frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		err := processLib(m)
		if err != nil {
			println("Error", err.Error())
			return
		}
	},
}
