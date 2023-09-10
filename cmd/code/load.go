package code

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"xs/internal/configs"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

var cmdLoad = &cobra.Command{
	Use:   "load",
	Short: "Tags section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		packet := args[0]
		pinning := wrappers.NewPinning()

		repo, err := pinning.FindRepo(packet)
		if err != nil {
			log.Fatal(err)
		}
		for name, _ := range *repo {
			println(name)
		}
		confirm := tools.ConfirmDialog("Load packets?")

		if confirm {
			fmt.Println("Proceeding with the action.")
			for name, val := range *repo {
				exec(val.Cid, val.To, name, pinning)
				//configs.LoadWorkspace() // todo random from config
			}
		} else {
			fmt.Println("Canceled.")
		}

	},
}

func exec(cid string, directory string, packageName string, pinning *wrappers.Pinning) {

	subDir := strings.Split(directory, "/")[0]

	subDirExists := tools.Exists(subDir)

	wsManager, err := configs.NewWsManager()
	if err != nil {
		log.Fatal(err)
	}

	sourceLoader := configs.NewSourceLoader()

	if !subDirExists {

		templateModule := wsManager.GetTemplateDirectory(subDir)
		println("templateModule", templateModule)

		repo, err := pinning.FindOne(templateModule)
		if err != nil {
			log.Fatal(err)
		}

		sourceLoader.Load(repo.Cid, subDir) // todo random from config
	}

	packPath := strings.Replace(packageName, "@", "/", -1)
	moduleSubDir := directory + packPath
	moduleSubDirExists := tools.Exists(moduleSubDir)
	if !moduleSubDirExists {
		sourceLoader.Load(cid, moduleSubDir)
	} else {
		println("already loaded", moduleSubDir)
	}

	// load package
	// add package
	// save ws file
	//repository := manager.GetSectionRepository(sectionName)
	//err := tools.CreateDirs(sectionName)
	//if err != nil {
	//	io.Panic(err)
	//}
	//pt := tools.NewPathTools()
	//pt.MoveTo(sectionName)
	//configs.LoadBase(repository)
	//pt.MoveToBasePath()
	//manager.SetSectionState(sectionName, "enabled")
	//manager.Save()

}
