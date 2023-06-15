package compilers

import (
	"os"
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
)

type Container struct {
	PrintConsole bool
}

func (n Container) Compile(params map[string]string) error {
	groupDir := "modules"

	m := params["name"]

	mod, extractError := configs.ExtractModule(m, groupDir, configs.BACK)
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

	arg := "build --platform=" + platform + "  --progress=plain --output type=image,name=" + reg + "/" + mod.Directory + ":latest,push=true ."
	io.Println(command + " " + arg)
	argsSplit := strings.Split(arg, " ")
	if errDir != nil {
		return errDir
	}

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "nerdctl", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	if result == 0 {
		io.PrintColor("OK", io.Green)
	} else {
		io.PrintColor("ERROR", io.Red)
	}

	return nil
}
