package jobs_build

import (
	"os"
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
)

type BuildContainer struct {
	params       map[string]string
	registry     string
	platform     string
	printConsole bool
}

func (c *BuildContainer) Execute() *jobs.Result {
	path := c.params["path"]
	name := c.params["name"]

	err := os.Chdir(path)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	command := "nerdctl"
	io.Println("command:" + command)

	pn := strings.Replace(name, "@", "/", 1)
	tag := c.registry + pn + ":latest"
	//arg := "build --platform=" + c.platform + "  --progress=plain --output type=image,name=" + tag + " ."
	arg := "build" + " -t " + tag + " --platform=" + c.platform + "  --progress=plain" + " ."
	io.Println(command + " " + arg)
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Key: c.Title().Key, Command: command, Args: argsSplit, PrintToConsole: c.printConsole}

	result := stdPrinter.Start()

	if result == 0 {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "BuildContainer executed",
		}
	} else {
		return &jobs.Result{
			Success:     false,
			Error:       nil,
			Description: "BuildContainer not executed",
		}
	}
}

func (b *BuildContainer) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: b.params["path"],
		Name:        b.params["name"],
		Key:         "build-container-" + b.params["name"],
	}
}

func NewBuildContainer(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &BuildContainer{
		params:       params,
		platform:     "amd64",
		registry:     "registry.solenopsys.org", // todo move to config
		printConsole: printConsole,
	}
}
