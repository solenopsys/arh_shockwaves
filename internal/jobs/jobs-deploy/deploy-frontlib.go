package jobs_deploy

import (
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type DeployFrontlib struct {
	params       map[string]string
	printConsole bool
}

func (d *DeployFrontlib) Execute() *jobs.Result {
	dest := d.params["dist"]
	src := d.params["path"]
	command := "pnpm"
	pt := xstool.PathTools{}
	pt.SetBasePathPwd() // remove it
	pt.MoveTo(src)
	absoluteDestPath, err := filepath.Abs(dest)
	if err != nil {
		io.Panic(err)
	}
	stdPrinter := io.StdPrinter{Key: d.Title().Key, Command: command, Args: []string{"publish", absoluteDestPath, "--no-git-checks"}, PrintToConsole: d.printConsole}
	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Deploy FrontLib executed",
		}
	} else {
		return &jobs.Result{
			Success:     false,
			Error:       nil,
			Description: "Deploy FrontLib not executed",
		}
	}
}

func (d *DeployFrontlib) Title() jobs.ItemTitle {
	name := d.params["name"]
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Deploy frontlib: " + name,
		Name:        name,
		Key:         "deploy-frontlib-" + name,
	}
}

func NewDeployFrontLib(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployFrontlib{
		params:       params,
		printConsole: printConsole,
	}
}
