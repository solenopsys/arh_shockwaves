package build

import (
	"github.com/spf13/cobra"
	"xs/services"
	"xs/utils"
)

func processLib(m string) error {
	groupDir := "packages"

	if m == "*" {
		println("SetCompiled all libraries")
		cc := services.NewLibCompileController("./xs.json", groupDir)
		println("Scan directories")
		cc.LoadPlan()
		println("Start compile")
		cc.CompileOnOneThread(false)
	} else {
		mod, extractError := services.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			return extractError
		}

		compiler := services.NpmCompileExecutor{PrintConsole: true}
		println("Mod ", mod.Directory)

		path := "./" + groupDir + "/" + mod.Directory
		println("SetCompiled library", path)

		dest := utils.LoadNgDest(path)
		compiler.Compile(path, dest)

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
