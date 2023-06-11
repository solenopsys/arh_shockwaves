package compilers

import (
	"errors"
	"os/exec"
	"strings"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

const NPM_APPLICATION = "pnpm"

type AngularPackageCompileExecutor struct {
	PrintConsole bool
}

func (n AngularPackageCompileExecutor) Compile(params map[string]string) error {
	src := params["path"]
	dest := params["dest"]
	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(src)
	arg := "build"
	argsSplit := strings.Split(arg, " ")
	stdPrinter := io.StdPrinter{Out: make(chan string), Command: NPM_APPLICATION, Args: argsSplit, PrintToConsole: n.PrintConsole}
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
		if result != 0 {
			io.PrintColor("ERROR PNPM LINK:"+string(rune(linkRes)), io.Red)
			return errors.New("ERROR PNPM LINK")
		}
		return nil

	} else {
		io.PrintColor("ERROR", io.Red)

		return errors.New("ERROR PNPM BUILD")
	}

}
