package build

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	services2 "xs/internal/services"
	"xs/pkg/wrappers"
)

func processLib(m string) error {
	groupDir := "packages"

	if m == "*" {
		println("SetCompiled all libraries")
		cc := services2.NewLibCompileController("./xs.json", groupDir)
		println("Scan directories")
		cc.LoadPlan()
		println("Start compile")
		cc.CompileOnOneThread(false)
	} else {
		mod, extractError := configs.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			return extractError
		}

		compiler := services2.NpmCompileExecutor{PrintConsole: true}
		println("Mod ", mod.Directory)

		path := "./" + groupDir + "/" + mod.Directory
		println("SetCompiled library", path)

		dest := wrappers.LoadNgDest(path)
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
