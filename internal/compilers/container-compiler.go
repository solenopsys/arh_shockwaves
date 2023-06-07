package compilers

import (
	"os"
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
)

type ContainerCompileExecutor struct {
	PrintConsole bool
}

func (n ContainerCompileExecutor) Compile(params map[string]string) error {
	groupDir := "modules"

	m := params["name"]

	mod, extractError := configs.ExtractModule(m, groupDir, "back")
	if extractError != nil {
		return extractError
	}
	path := "./" + groupDir + "/" + mod.Directory

	platform := "amd64"
	reg := "registry.solenopsys.org"

	errDir := os.Chdir(path)
	if errDir != nil {
		return errDir
	}

	command := "nerdctl"
	io.Println("command:" + command)

	arg := "build --platform=" + platform + " --output type=image,name=" + reg + "/" + mod.Directory + ":latest,push=true ."
	argsSplit := strings.Split(arg, " ")
	if errDir != nil {
		return errDir
	}

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "nerdctl", Args: argsSplit}
	go stdPrinter.Processing()
	stdPrinter.Start()

	return nil
}
