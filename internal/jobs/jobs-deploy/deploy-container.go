package jobs_deploy

import (
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployContainer struct {
	params       map[string]string
	registry     string
	printConsole bool
}

func (d *DeployContainer) Execute() *jobs.Result {
	name := d.params["name"]

	command := "nerdctl"

	arg := " push " + d.registry + "/" + name + ":latest "
	io.Println(command + " " + arg)
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: command, Args: argsSplit, PrintToConsole: d.printConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	if result == 0 {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Deploy Container executed",
		}
	} else {
		return &jobs.Result{
			Success:     false,
			Error:       nil,
			Description: "Deploy Container not executed",
		}
	}
}

func (d *DeployContainer) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Deploy container: " + d.params["name"],
		Short:       "Reddy",
	}
}

func NewDeployContainer(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployContainer{
		registry:     "registry.solenopsys.org", // todo move to config
		params:       params,
		printConsole: printConsole,
	}
}
