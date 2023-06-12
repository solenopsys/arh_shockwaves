package services

import (
	"xs/internal"
	"xs/internal/configs"
	"xs/pkg/io"
)

type UniversalCompileController struct {
	Executor  internal.CompileExecutor
	Extractor internal.CompileParamsExtractor
	GroupDir  string
	libs      []*configs.XsMonorepoModule
	xm        *configs.XsManager
	RepoType  string
}

func (u *UniversalCompileController) SelectLibs(name string) error {
	u.xm = &configs.XsManager{}
	err := u.xm.Load("./xs-treerepo.json")
	if err != nil {
		io.Panic(err)
	}

	libs := u.xm.FilterLibs(name, u.GroupDir)

	for _, lib := range libs {
		io.Print(lib.Name, " ")
	}

	u.libs = libs
	return err
}

func (u *UniversalCompileController) CompileSelectedLibs() {
	if len(u.libs) == 1 {
		u.compileOne(u.libs[0].Name)
	} else if len(u.libs) > 1 {
		u.compileMany()
	}
}

func (u *UniversalCompileController) compileMany() {
	if u.RepoType == configs.FRONT {
		io.Println("SetCompiled all libraries")
		//	executor.PrintConsole= false
		libCompiler := NewLibCompileController(u.xm, u.Executor)
		io.Println("Scan directories")
		libCompiler.LoadPlan(u.GroupDir, u.libs)
		io.Println("Start compile")
		libCompiler.CompileOnOneThread(false, u.GroupDir, u.Extractor)
	} else {
		for _, lib := range u.libs {
			u.compileOne(lib.Name)
		}
	}
}

func (u *UniversalCompileController) compileOne(name string) {

	mod, extractError := configs.ExtractModule(name, u.GroupDir, u.RepoType)
	if extractError != nil {
		io.Panic(extractError)
	}

	compiler := u.Executor
	if mod == nil {
		io.Panic("Module not found")
	}
	io.Println("Mod ", mod.Directory)

	path := "./" + u.GroupDir + "/" + mod.Directory
	io.Println("SetCompiled library", path)

	var params = map[string]string{}

	if u.Extractor != nil {
		params = u.Extractor.Extract(name, path)
	}

	err := compiler.Compile(params)

	if err != nil {
		io.Panic(err)
	}
}
