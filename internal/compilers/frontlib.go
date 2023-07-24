package compilers

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

const NPM_APPLICATION = "pnpm"

type Frontlib struct {
	PrintConsole bool
}

func publishOne(dist string) error {
	io.Println("Publish library", dist)
	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(dist)

	println("CURRENT DIR", pt.GetPwd())

	args := []string{"publish", "--no-git-checks", "--access", "public"}

	cmd := exec.Command(NPM_APPLICATION, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		io.Println(err)
	}

	// Set the Stdout of the command to our buffer

	linkRes := cmd.ProcessState.ExitCode()

	pt.MoveToBasePath()
	if linkRes != 0 {
		o := stdout.String()

		io.Println("ERROR PNPM PUBLISH: " + o)
		return errors.New("ERROR PNPM PUBLISH")
	}

	return nil
}

func (n Frontlib) Compile(params map[string]string) error {
	src := params["path"]
	dest := params["dest"]
	publish := params["publish"]
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

		const PUBLISH = "publish"

		if publish == "true" {
			err := publishOne(params["dest"])
			if err != nil {
				io.Println(err)
			}
		}
		return nil

	} else {
		io.PrintColor("ERROR", io.Red)

		return errors.New("ERROR PNPM BUILD")
	}

}
