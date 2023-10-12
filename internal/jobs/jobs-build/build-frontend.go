package jobs_build

import (
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type BuildFrontend struct {
	printConsole bool
	params       map[string]string
}

func NewBuildFrontend(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &BuildFrontend{printConsole, params}
}

func (b *BuildFrontend) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: b.params["path"],
		Name:        b.params["name"],
	}
}

func (n *BuildFrontend) Execute() *jobs.Result {

	pt := xstool.PathTools{}
	src := n.params["path"]

	pt.SetBasePathPwd()
	pt.MoveTo(src)

	arg := "build"
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "ng", Args: argsSplit, PrintToConsole: n.printConsole}
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
