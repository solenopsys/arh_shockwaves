package services

import (
	"xs/internal"
	"xs/internal/configs"
	"xs/pkg/io"
)

func CompileGroup(name string, groupDir string, executor internal.CompileExecutor, extractor internal.CompileParamsExtractor) {

	xm := &configs.XsManager{}

	err := xm.Load("./xs-treerepo.json")
	if err != nil {
		io.Panic(err)
	}

	libs := xm.FilterLibs(name, groupDir)

	for _, lib := range libs {
		io.Print(lib.Name, " ")
	}

	if len(libs) == 1 {
		mod, extractError := configs.ExtractModule(name, groupDir, "front")
		if extractError != nil {
			err = extractError
		}

		compiler := executor
		io.Println("Mod ", mod.Directory)

		path := "./" + groupDir + "/" + mod.Directory
		io.Println("SetCompiled library", path)

		var params = map[string]string{}

		if extractor != nil {
			params = extractor.Extract(name, path)
		}

		compiler.Compile(params)

	} else if len(libs) > 1 {
		io.Println("SetCompiled all libraries")
		//	executor.PrintConsole= false
		cc := NewLibCompileController(xm, executor)
		io.Println("Scan directories")
		cc.LoadPlan(groupDir, libs)
		io.Println("Start compile")
		cc.CompileOnOneThread(false, groupDir, extractor)
	}

	if err != nil {
		io.Println("Error", err.Error())
		return
	}
}
