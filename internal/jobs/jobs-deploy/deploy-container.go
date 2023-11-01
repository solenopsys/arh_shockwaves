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
	pn := strings.Replace(name, "@", "/", 1)
	arg := "push " + d.registry + pn + ":latest"
	io.Println(command + " " + arg)
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Key: d.Title().Key, Command: command, Args: argsSplit, PrintToConsole: d.printConsole}
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

func (d *DeployContainer) Title() jobs.ItemTitle {
	name := d.params["name"]
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Deploy container: " + name,
		Name:        name,
		Key:         "deploy-container-" + name,
	}
}

func NewDeployContainer(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployContainer{
		registry:     "registry.solenopsys.org", // todo move to config
		params:       params,
		printConsole: printConsole,
	}
}
