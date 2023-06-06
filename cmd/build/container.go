package build

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"xs/internal/configs"
	services2 "xs/pkg/tools"
)

var cmdContainer = &cobra.Command{
	Use:   "container [name]",
	Short: "Containers for module build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		m := args[0]
		groupDir := "modules"

		mod, extractError := configs.ExtractModule(m, groupDir, "back")
		if extractError != nil {
			println("Error", extractError.Error())
			return
		}
		path := "./" + groupDir + "/" + mod.Directory

		platform := "amd64"
		reg := "registry.solenopsys.org"

		errDir := os.Chdir(path)
		if errDir != nil {
			fmt.Println(errDir)
			return
		}

		command := "nerdctl"
		println("command:" + command)

		arg := "build --platform=" + platform + " --output type=image,name=" + reg + "/" + mod.Directory + ":latest,push=true ."
		argsSplit := strings.Split(arg, " ")
		if errDir != nil {
			fmt.Println(errDir)
			return
		}

		stdPrinter := services2.StdPrinter{Out: make(chan string), Command: "nerdctl", Args: argsSplit}
		go stdPrinter.Processing()
		stdPrinter.Start()
	},
}
