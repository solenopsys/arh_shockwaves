package build

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	services2 "xs/internal/services"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

func processLib(m string) error {
	groupDir := "packages"

	if m == "*" {
		io.Println("SetCompiled all libraries")
		cc := services2.NewLibCompileController("./xs.json", groupDir)
		io.Println("Scan directories")
		cc.LoadPlan()
		io.Println("Start compile")
		cc.CompileOnOneThread(false)
	} else {
		mod, extractError := configs.ExtractModule(m, groupDir, "front")
		if extractError != nil {
			return extractError
		}

		compiler := services2.NpmCompileExecutor{PrintConsole: true}
		io.Println("Mod ", mod.Directory)

		path := "./" + groupDir + "/" + mod.Directory
		io.Println("SetCompiled library", path)

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
			io.Println("Error", err.Error())
			return
		}
	},
}
