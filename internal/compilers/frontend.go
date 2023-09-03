package compilers

import (
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type Frontend struct {
	PrintConsole bool
}

func (n Frontend) Compile(params map[string]string) error {
	groupDir := "endpoints"

	pt := xstool.PathTools{}
	src := params["path"]
	m := params["name"]

	_, extractError := configs.ExtractModule(m, groupDir, configs.FRONT)
	if extractError != nil {
		return extractError
	}
	pt.SetBasePathPwd()
	pt.MoveTo(src)

	arg := "build"
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "ng", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		io.PrintColor("OK", io.Green)
	} else {
		io.PrintColor("ERROR", io.Red)
	}

	return nil
}
