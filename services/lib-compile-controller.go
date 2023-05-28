package services

import (
	"encoding/json"
	"github.com/fatih/color"
	"os/exec"
	"strings"
	"xs/utils"
)

type BuildConfig struct {
	Dest string `json:"dest"`
}

type CompileCommand struct {
	LibName      string
	LibDirectory string
}

type CompileExecutor interface {
	Compile(path string)
}

type NpmCompileExecutor struct {
	PrintConsole bool
}

func (n NpmCompileExecutor) loadDest() string {
	bc := &BuildConfig{}
	bytesFromFile, err := utils.ReadFile("ng-package.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), bc)
	if err != nil {
		panic(err)
	}

	return bc.Dest
}

func (n NpmCompileExecutor) Compile(path string) {
	pt := utils.PathTools{}
	pt.SetBasePathPwd()

	//

	pt.MoveTo(path)
	arg := "build"
	argsSplit := strings.Split(arg, " ")
	stdPrinter := StdPrinter{Out: make(chan string), Command: "pnpm", Args: argsSplit, PrintToConsole: n.PrintConsole}
	go stdPrinter.Processing()
	result := stdPrinter.Start()

	destPath := n.loadDest()
	destFixed := strings.Replace(destPath, "../../../", "./", -1)
	pt.MoveToBasePath()

	if result == 0 {
		c := color.New(color.BgHiGreen, color.Bold)
		c.Print(" OK ")
		println("")

		println("Make link: ", destFixed)
		cmd := exec.Command("pnpm", "link", destFixed)

		if err := cmd.Start(); err != nil {
			panic(err)
		}
		cmd.Wait()
		linkRes := cmd.ProcessState.ExitCode()
		if result != 0 {
			color.Red("ERROR PNPM LINK ", linkRes)
		}
	} else {
		c := color.New(color.BgHiRed, color.Bold)
		c.Print(" ERROR ")
		println("")

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

func (c *LibCompileController) CompileOnOneThread() {

	println("Packages count: ", c.packagesOrder.count())
	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}
		println("COMPILE GROUP: ")
		for _, pack := range list {
			println("\t", pack.Name+" ")
		}

		for _, pack := range list {
			xsPackConf := c.xsManager.extract(c.libGroup, pack.Name)

			if xsPackConf == nil {
				panic(pack.Name + " not found in " + c.xsFile + " group " + c.libGroup)
			}

			path := c.libGroup + "/" + xsPackConf.Directory
			println("COMPILE: "+path, " --------->")
			c.compileNow[pack.Name] = true
			c.compileExecutor.Compile(path)
			c.compileNow[pack.Name] = false
			c.packagesOrder.Compile(pack.Name)
		}

	}
}
