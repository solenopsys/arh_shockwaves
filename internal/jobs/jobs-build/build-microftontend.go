package jobs_build

import (
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
)

type MicroFronted struct {
	PrintConsole bool
	params       map[string]string
}

func (b *MicroFronted) Execute() *jobs.Result {
	//	groupDir := "modules"

	lib := b.params["lib"]
	//	m := params["name"]

	arg := "bmf " + lib
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit, PrintToConsole: b.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	if result == 0 {
		io.PrintColor("OK", io.Green)
	} else {
		io.PrintColor("ERROR", io.Red)
	}

	return nil
}

func (b *MicroFronted) Description() string {
	return "Build helm " + b.params["lib"]
}

func NewMicroFronted(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &MicroFronted{
		params: params,
	}
}
