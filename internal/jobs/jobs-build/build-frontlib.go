package jobs_build

import (
	"errors"
	"os/exec"
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

const NPM_APPLICATION = "pnpm"

type BuildFrontLib struct {
	params       map[string]string
	printConsole bool
}

func (b *BuildFrontLib) Execute() *jobs.Result {
	src := b.params["path"]
	dest := b.params["dest"]
	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(src)
	arg := "build"
	argsSplit := strings.Split(arg, " ")
	stdPrinter := io.StdPrinter{Out: make(chan string), Command: NPM_APPLICATION, Args: argsSplit, PrintToConsole: b.printConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		io.PrintColor("OK", io.Green)

		//io.Println("Make link: ", dest)
		cmd := exec.Command(NPM_APPLICATION, "link", dest)

		if err := cmd.Start(); err != nil {
			io.Panic(err)
		}
		cmd.Wait()
		linkRes := cmd.ProcessState.ExitCode()
		if linkRes != 0 {
			return &jobs.Result{
				Success: false,
				Err:     errors.New("ERROR PNPM LINK"),
			}
		}

		return &jobs.Result{
			Success:     true,
			Err:         nil,
			Description: "BuildLib executed",
		}

	} else {

		return &jobs.Result{
			Success: false,
			Err:     errors.New("ERROR PNPM BUILD"),
		}

	}

}
