package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployFrontlib struct {
	params       map[string]string
	printConsole bool
}

func (d *DeployFrontlib) Execute() *jobs.Result {
	dest := d.params["dist"]
	command := "pnpm"
	stdPrinter := io.StdPrinter{Key: d.Title().Name, Command: command, Args: []string{"publish", dest}, PrintToConsole: d.printConsole}
	result := stdPrinter.Start()

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
	}
}

func NewDeployFrontLib(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployFrontlib{
		params:       params,
		printConsole: printConsole,
	}
}
