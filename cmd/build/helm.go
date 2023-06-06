package build

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		groupDir := "modules"
		mod, extractError := configs.ExtractModule(m, groupDir, "back")
		if extractError != nil {
			io.Println("Error", extractError.Error())
			return
		}
		path := "./" + groupDir + "/" + mod.Directory + "/install"

		io.Println("path", path)
		arch := wrappers.ArchiveDir(path, m)

		// write archive to file
		io.Println("archive size", len(arch))

		wrappers.PushDir(arch)
	},
}
