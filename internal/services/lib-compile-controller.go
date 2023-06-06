package services

import (
	"errors"
	"os/exec"
	"strings"
	"xs/internal/configs"
	"xs/pkg/io"
	tools2 "xs/pkg/tools"
	"xs/pkg/wrappers"
)

type CompileCommand struct {
	LibName      string
	LibDirectory string
}

type CompileExecutor interface {
	Compile(path string, dest string) error
}

type NpmCompileExecutor struct {
	PrintConsole bool
}

func (n NpmCompileExecutor) Compile(src string, dest string) error {
	pt := tools2.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(src)
	arg := "build"
	argsSplit := strings.Split(arg, " ")
	stdPrinter := io.StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		io.PrintColor("OK", io.Green)

		//io.Println("Make link: ", dest)
		cmd := exec.Command("pnpm", "link", dest)

		if err := cmd.Start(); err != nil {
			panic(err)
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

func NewLibCompileController(xsFile string, libGroup string) *LibCompileController {
	c := &LibCompileController{}
	c.xsFile = xsFile
	c.libGroup = libGroup
	c.compileNow = map[string]bool{}
	c.compileExecutor = NpmCompileExecutor{PrintConsole: false}
	c.LoadPlan()
	return c
}

type LibCompileController struct {
	xsFile          string
	libGroup        string
	compileNow      map[string]bool
	packagesOrder   *NpmLibPackagesOrder
	compileExecutor CompileExecutor
	xsManager       *configs.XsManager
}

func (c *LibCompileController) LoadPlan() {
	fileName := c.xsFile
	xm := &configs.XsManager{}
	c.xsManager = xm
	err := xm.Load(fileName)
	if err != nil {
		panic(err)
	}

	libs := xm.ExtractGroup(c.libGroup)
	ord := NewNpmLibPackagesOrder(false)

	for _, lib := range libs {
		lp := configs.LoadNpmLibPackage("./" + c.libGroup + "/" + lib.Directory + "/package.json")
		ord.AddPackage(lp)
	}

	c.packagesOrder = ord
}

func (c *LibCompileController) CompileOnOneThread(force bool) { //todo need refactoring

	cache := NewCompileCache(".xs/compiled")

	io.Println("Packages count: ", c.packagesOrder.count())
	var n = 0
	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}
		//io.Println("COMPILE GROUP: ")
		//for _, pack := range list {
		//	io.Println("\t", pack.Name+" ")
		//}

		for _, pack := range list {
			xsPackConf := c.xsManager.Extract(c.libGroup, pack.Name)

			if xsPackConf == nil {
				panic(pack.Name + " not found in " + c.xsFile + " group " + c.libGroup)
			}

			path := c.libGroup + "/" + xsPackConf.Directory
			n++
			io.PrintColor(string(rune(n))+" : "+xsPackConf.Npm+" ", io.Blue)
			c.compileNow[pack.Name] = true

			dest := wrappers.LoadNgDest(path)

			dirExists := tools2.DirExists(dest)
			excludeDirs := []string{"node_modules"}
			var hashesOk = false
			if dirExists { //todo refactor

				srcHash, errHash := tools2.HashOfDir(path, excludeDirs)
				if errHash != nil {
					panic(errHash)
				}
				dstHash, errHash := tools2.HashOfDir(dest, excludeDirs)
				if errHash != nil {
					panic(errHash)
				}
				hashesOk = cache.checkHash(srcHash, dstHash)
			}
			if hashesOk && !force {
				io.PrintColor("SKIP", io.Blue)

				c.packagesOrder.SetCompiled(pack.Name)
				c.compileNow[pack.Name] = false
			} else {
				err := c.compileExecutor.Compile(path, dest)
				if err != nil {
					panic(err)
				} else {
					srcHash, errHash := tools2.HashOfDir(path, excludeDirs)
					if errHash != nil {
						panic(errHash)
					}
					dstHash, errHash := tools2.HashOfDir(dest, excludeDirs)
					if errHash != nil {
						panic(errHash)
					}
					errHash = cache.saveHash(srcHash, dstHash)
					if errHash != nil {
						panic(errHash)
					}
					c.compileNow[pack.Name] = false
					c.packagesOrder.SetCompiled(pack.Name)
				}
			}

		}

	}
}
