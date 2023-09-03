package compilers

import (
	"strings"
	"xs/pkg/io"
)

type Microfronted struct {
	PrintConsole bool
}

func (n Microfronted) Compile(params map[string]string) error {
	//	groupDir := "modules"

	lib := params["lib"]
	//	m := params["name"]

	arg := "bmf " + lib
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	if result == 0 {
		io.PrintColor("OK", io.Green)
	} else {
		io.PrintColor("ERROR", io.Red)
	}

	return nil
}
