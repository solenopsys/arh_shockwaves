package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/configs"
	services2 "xs/internal/services"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdFrontlib = &cobra.Command{
	Use:   "frontlib [name]",
	Short: "Frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		groupDir := "packages"

		var err error

		if name == "*" {
			io.Println("SetCompiled all libraries")
			cc := services2.NewLibCompileController("./xs.json", groupDir)
			io.Println("Scan directories")
			cc.LoadPlan()
			io.Println("Start compile")
			cc.CompileOnOneThread(false)
		} else {
			mod, extractError := configs.ExtractModule(name, groupDir, "front")
			if extractError != nil {
				err = extractError
			}

			compiler := compilers.AngularPackageCompileExecutor{PrintConsole: true}
			io.Println("Mod ", mod.Directory)

			path := "./" + groupDir + "/" + mod.Directory
			io.Println("SetCompiled library", path)

			dest := wrappers.LoadNgDest(path)
			params := map[string]string{
				"path": path,
				"dest": dest,
			}
			compiler.Compile(params)
		}

		if err != nil {
			io.Println("Error", err.Error())
			return
		}
	},
}
