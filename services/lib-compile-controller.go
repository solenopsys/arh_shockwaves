package services

import (
	"errors"
	"github.com/fatih/color"
	"os/exec"
	"strings"
	"xs/utils"
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
	pt := utils.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(src)
	arg := "build"
	argsSplit := strings.Split(arg, " ")
	stdPrinter := StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	pt.MoveToBasePath()

	if result == 0 {
		c := color.New(color.BgHiGreen, color.Bold)
		c.Print(" OK ")
		println("")

		//println("Make link: ", dest)
		cmd := exec.Command("pnpm", "link", dest)

		if err := cmd.Start(); err != nil {
			panic(err)
		}
		cmd.Wait()
		linkRes := cmd.ProcessState.ExitCode()
		if result != 0 {
			color.Red("ERROR PNPM LINK ", linkRes)
			return errors.New("ERROR PNPM LINK")
		}
		return nil

	} else {
		c := color.New(color.BgHiRed, color.Bold)
		c.Print(" ERROR ")
		println("")

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
	xsManager       *XsManager
}

func (c *LibCompileController) LoadPlan() {
	fileName := c.xsFile
	xm := &XsManager{}
	c.xsManager = xm
	err := xm.load(fileName)
	if err != nil {
		panic(err)
	}

	libs := xm.extractGroup(c.libGroup)
	ord := NewNpmLibPackagesOrder(false)

	for _, lib := range libs {
		lp := utils.LoadNpmLibPackage("./" + c.libGroup + "/" + lib.Directory + "/package.json")
		ord.AddPackage(lp)
	}

	c.packagesOrder = ord
}

func (c *LibCompileController) CompileOnOneThread(force bool) { //todo need refactoring

	cache := NewCompileCache(".xs/compiled")

	println("Packages count: ", c.packagesOrder.count())
	var n = 0
	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}
		//println("COMPILE GROUP: ")
		//for _, pack := range list {
		//	println("\t", pack.Name+" ")
		//}

		for _, pack := range list {
			xsPackConf := c.xsManager.extract(c.libGroup, pack.Name)

			if xsPackConf == nil {
				panic(pack.Name + " not found in " + c.xsFile + " group " + c.libGroup)
			}

			path := c.libGroup + "/" + xsPackConf.Directory
			n++
			print(n, " COMPILE: "+xsPackConf.Npm+" ")
			c.compileNow[pack.Name] = true

			dest := utils.LoadNgDest(path)

			dirExists := utils.DirExists(dest)
			excludeDirs := []string{"node_modules"}
			var hashesOk = false
			if dirExists { //todo refactor

				srcHash, errHash := utils.HashOfDir(path, excludeDirs)
				if errHash != nil {
					panic(errHash)
				}
				dstHash, errHash := utils.HashOfDir(dest, excludeDirs)
				if errHash != nil {
					panic(errHash)
				}
				hashesOk = cache.checkHash(srcHash, dstHash)
			}
			if hashesOk && !force {
				color := color.New(color.BgHiBlue, color.Bold)
				color.Print(" SKIP ")
				println("")

				c.packagesOrder.SetCompiled(pack.Name)
				c.compileNow[pack.Name] = false
			} else {
				err := c.compileExecutor.Compile(path, dest)
				if err != nil {
					panic(err)
				} else {
					srcHash, errHash := utils.HashOfDir(path, excludeDirs)
					if errHash != nil {
						panic(errHash)
					}
					dstHash, errHash := utils.HashOfDir(dest, excludeDirs)
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
