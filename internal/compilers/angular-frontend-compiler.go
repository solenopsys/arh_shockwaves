package compilers

import (
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
)

type AngularFrontendCompileExecutor struct {
	PrintConsole bool
}

func (n AngularFrontendCompileExecutor) Compile(params map[string]string) error {
	groupDir := "modules"

	m := params["name"]

	mod, extractError := configs.ExtractModule(m, groupDir, "entrances")
	if extractError != nil {
		return extractError
	}

	arg := mod.Directory + "build:production"
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "ng", Args: argsSplit}
	go stdPrinter.Processing()
	stdPrinter.Start()

	return nil
}
