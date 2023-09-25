package jobs_build

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"xs/internal/compilers/sorters/fl"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

const NPM_APPLICATION = "pnpm"

type BuildFrontLib struct {
	params       map[string]string
	printConsole bool
}

func NewBuildFrontLib(params map[string]string, printConsole bool) *BuildFrontLib {
	return &BuildFrontLib{params, printConsole}
}

func (b *BuildFrontLib) saveToCache(dest string, path string, excludeDirs []string) {
	cache := fl.NewCompileCache(".xs/compiled")
	srcHash, errHash := xstool.HashOfDir(path, excludeDirs)
	if errHash != nil {
		io.Panic(errHash)
	}
	dstHash, errHash := xstool.HashOfDir(dest, excludeDirs)
	if errHash != nil {
		io.Panic(errHash)
	}
	errHash = cache.SaveHash(srcHash, dstHash)
	if errHash != nil {
		io.Panic(errHash)
	}

}

func (b *BuildFrontLib) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Build FrontLib " + b.params["path"],
		Short:       "Reddy",
	}
}

func (b *BuildFrontLib) Execute() *jobs.Result { // todo refactoring
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

	absoluteDestPath, err := filepath.Abs(dest)
	absoluteSrcPath, err := filepath.Abs(".")
	if err != nil {
		io.Panic(err)
	}
	pt.MoveToBasePath()

	if result == 0 {
		io.PrintColor("OK", io.Green)

		io.Println("Make link: ", absoluteDestPath)
		pt.MoveTo("frontends") //todo move to const
		currentDir, err := os.Getwd()

		println("current path", currentDir)

		cmd := exec.Command(NPM_APPLICATION, "link", absoluteDestPath)

		if err := cmd.Start(); err != nil {
			io.Panic(err)
		}
		err = cmd.Wait()
		if err != nil {
			io.Panic(err)
		}
		linkRes := cmd.ProcessState.ExitCode()
		if linkRes != 0 {
			return &jobs.Result{
				Success: false,
				Error:   errors.New("ERROR PNPM LINK"),
			}
		}

		b.saveToCache(absoluteDestPath, absoluteSrcPath, []string{"node_modules"})
		pt.MoveToBasePath()
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "BuildLib executed",
		}

	} else {
		pt.MoveToBasePath()
		return &jobs.Result{
			Success: false,
			Error:   errors.New("ERROR PNPM BUILD"),
		}

	}

}
