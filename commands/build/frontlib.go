package build

import (
	"github.com/spf13/cobra"
	"xs/services"
)

func processLib(m string) error {
	groupDir := "packages"

	if m == "*" {
		println("Compile all libraries")
		cc := services.NewLibCompileController("./xs.json", groupDir)
		println("Scan directories")
		cc.LoadPlan()
		println("Start compile")
		cc.CompileOnOneThread()
	} else {
		mod, extractError := services.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			return extractError
		}

		compiler := services.NpmCompileExecutor{PrintConsole: true}
		println("Mod ", mod.Directory)

		path := "./" + groupDir + "/" + mod.Directory
		println("Compile library", path)
		compiler.Compile(path)

	}

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
