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

func (b *MicroFronted) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Build helm " + b.params["lib"],
		Short:       "Reddy",
	}
}

func NewMicroFronted(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &MicroFronted{
		params: params,
	}
}
