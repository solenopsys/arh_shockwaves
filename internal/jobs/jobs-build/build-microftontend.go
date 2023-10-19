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

	stdPrinter := io.StdPrinter{Key: b.Title().Key, Command: NPM_APPLICATION, Args: argsSplit, PrintToConsole: b.PrintConsole}
	return stdPrinter.Start()
}

func (b *MicroFronted) Execute() *jobs.Result {

	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo("frontends") //todo move to const

	fl := services.NewFrontLibController(b.Title().Key)

	fl.PreProcessing()
	result := b.build()
	fl.PostProcessing()

	defer pt.MoveToBasePath()

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

func (b *MicroFronted) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: b.params["lib"],
		Name:        b.params["name"],
		Key:         "build-microfrontend-" + b.params["name"],
	}
}

func NewMicroFronted(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &MicroFronted{
		PrintConsole: printConsole,
		params:       params,
	}
}
