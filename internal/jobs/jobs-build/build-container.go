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

	arg := "build --platform=" + c.platform + "  --progress=plain --output type=image,name=" + c.registry + "/" + name + ":latest,push=true ."
	io.Println(command + " " + arg)
	argsSplit := strings.Split(arg, " ")

	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "nerdctl", Args: argsSplit, PrintToConsole: c.printConsole}
	go stdPrinter.Processing()
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

func NewBuildContainer(params map[string]string) *BuildContainer {
	return &BuildContainer{params: params, platform: "amd64", registry: "registry.solenopsys.org"}
}