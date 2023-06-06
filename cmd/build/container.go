package build

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
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
			io.Println("Error", extractError.Error())
			return
		}
		path := "./" + groupDir + "/" + mod.Directory

		platform := "amd64"
		reg := "registry.solenopsys.org"

		errDir := os.Chdir(path)
		if errDir != nil {
			io.Println(errDir)
			return
		}

		command := "nerdctl"
		io.Println("command:" + command)

		arg := "build --platform=" + platform + " --output type=image,name=" + reg + "/" + mod.Directory + ":latest,push=true ."
		argsSplit := strings.Split(arg, " ")
		if errDir != nil {
			io.Println(errDir)
			return
		}

		stdPrinter := io.StdPrinter{Out: make(chan string), Command: "nerdctl", Args: argsSplit}
		go stdPrinter.Processing()
		stdPrinter.Start()
	},
}
