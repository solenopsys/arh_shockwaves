package jobs_build

import (
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type BuildFrontend struct {
	PrintConsole bool
	params       map[string]string
}

func (n *BuildFrontend) Execute() *jobs.Result {

	pt := xstool.PathTools{}
	src := n.params["path"]

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