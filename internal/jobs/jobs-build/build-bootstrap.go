package jobs_build

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type BuildFrontend struct {
	printConsole bool
	params       map[string]string
}

func NewBuildFrontend(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &BuildFrontend{printConsole, params}
}

func (b *BuildFrontend) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: b.params["path"],
		Name:        b.params["name"],
		Key:         "build-bootstrap-" + b.params["name"],
	}
}

func (n *BuildFrontend) Execute() *jobs.Result {

	pt := xstool.PathTools{}
	src := n.params["path"]
	packageName := strings.Replace(n.params["name"], "@", "", 1)

	pt.SetBasePathPwd()
	pt.MoveTo(src)

	scriptName := viper.GetString("scriptsNames.buildBootstrap")

	stdPrinter := io.StdPrinter{Key: n.Title().Key, Command: "pnpm", Args: []string{scriptName, packageName}, PrintToConsole: n.printConsole}

	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Build bootstrap executed",
		}
	} else {
		return &jobs.Result{
			Success:     false,
			Error:       errors.New("Build bootstrap  failed"),
			Description: "Build bootstrap  failed",
		}
	}
}
