package jobs_build

import (
	"errors"
	"strings"
	"xs/internal/jobs"
	"xs/internal/services"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type MicroFronted struct {
	PrintConsole bool
	params       map[string]string
}

func (b *MicroFronted) build() int {
	lib := strings.Replace(b.params["lib"], "./frontends", ".", 1) // todo remove replace
	//	m := params["name"]

	arg := "bmf " + lib
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: NPM_APPLICATION, Args: argsSplit, PrintToConsole: b.PrintConsole}
	go stdPrinter.Processing()
	return stdPrinter.Start()
}

func (b *MicroFronted) Execute() *jobs.Result {

	pt := xstool.PathTools{}
	pt.MoveTo("frontends") //todo move to const
	pt.SetBasePathPwd()
	defer pt.MoveToBasePath()

	fl := services.NewFrontLibController()

	fl.PreProcessing()
	result := b.build()
	fl.PostProcessing()

	if result == 0 {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Build microfrontend executed",
		}
	} else {
		return &jobs.Result{
			Success:     false,
			Error:       errors.New("Build microfrontend  failed"),
			Description: "Build microfrontend  failed",
		}
	}

}

func (b *MicroFronted) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Build microfrontend " + b.params["lib"],
		Short:       "Reddy",
	}
}

func NewMicroFronted(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &MicroFronted{
		PrintConsole: printConsole,
		params:       params,
	}
}
