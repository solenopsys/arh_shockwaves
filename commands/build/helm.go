package build

import (
	"github.com/spf13/cobra"
	"xs/services"
	"xs/utils"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		groupDir := "modules"
		mod, extractError := services.ExtractModule(m, groupDir)
		if extractError != nil {
			println("Error", extractError.Error())
			return
		}
		path := "./" + groupDir + "/" + mod.Directory + "/install"

		println("path", path)
		arch := utils.ArchiveDir(path, m)

		// write archive to file
		println("archive size", len(arch))

		utils.PushDir(arch)
	},
}
