package jobs_build

import (
	"errors"
	"os/exec"
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
	cache        *fl.CompileCache
}

func NewBuildFrontLib(params map[string]string, printConsole bool) *BuildFrontLib {
	cache := fl.NewCompileCache(".xs/compiled")
	return &BuildFrontLib{params, printConsole, cache}
}

func (b *BuildFrontLib) saveToCache(dest string, path string, excludeDirs []string) {
	srcHash, errHash := xstool.HashOfDir(path, excludeDirs)
	if errHash != nil {
		io.Panic(errHash)
	}
	dstHash, errHash := xstool.HashOfDir(dest, excludeDirs)
	if errHash != nil {
		io.Panic(errHash)
	}
	errHash = b.cache.SaveHash(srcHash, dstHash)
	if errHash != nil {
		io.Panic(errHash)
	}

}

func (b *BuildFrontLib) Description() string {
	return "BuildFrontLib " + b.params["name"]
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
				Error:   errors.New("ERROR PNPM LINK"),
			}
		}

		b.saveToCache(dest, src, []string{"node_modules"})

		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "BuildLib executed",
		}

	} else {

		return &jobs.Result{
			Success: false,
			Error:   errors.New("ERROR PNPM BUILD"),
		}

	}

}
