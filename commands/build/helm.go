package build

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/os"
	"io/ioutil"
	"xs/services"
	"xs/utils"
)

var cmdHelm = &cobra.Command{
	Use:   "helm [name]",
	Short: "Helm build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := args[0]
		fileName := "./xs.json"
		exists := os.FileExists(fileName)
		if exists {
			config := services.LoadConfigFile(fileName)
			repoType := services.FileTypeMapping[config.Format.Name]

			if repoType == "front" {
				println("Not applicable for front monorepo")
			} else if repoType == "back" {
				groupDir := "modules"
				groups := config.Groups
				modules := groups[groupDir]
				var okModule services.XsMonorepoModule
				var ok = false
				for _, module := range modules {
					if module.Directory == m {
						okModule = module
						ok = true
					}
				}

				if ok {
					println("Ok")

					path := "./" + groupDir + "/" + okModule.Directory + "/install"

					arch := utils.ArchiveDir(path, m)

					// write archive to file
					ioutil.WriteFile("bla-bla-0.1.8.tgz", arch, 0644)
					println("archive size", len(arch))

					utils.PushDir(arch)
				} else {
					println("Module not found")
				}
			} else {
				println("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
				return
			}

		} else {
			log.Error("xs.json not found, directory not initialized")
		}
	},
}
