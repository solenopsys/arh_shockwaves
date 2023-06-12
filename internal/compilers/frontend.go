package compilers

import (
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
)

type Frontend struct {
	PrintConsole bool
}

func (n Frontend) Compile(params map[string]string) error {
	groupDir := "modules"

	m := params["name"]

	_, extractError := configs.ExtractModule(m, groupDir, "entrances")
	if extractError != nil {
		return extractError
	}

	arg := "build:production"
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "ng", Args: argsSplit}
	go stdPrinter.Processing()
	stdPrinter.Start()

	return nil
}
