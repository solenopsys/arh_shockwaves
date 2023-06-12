package compilers

import (
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type Helm struct {
	PrintConsole bool
}

func (n Helm) Compile(params map[string]string) error {
	m := params["name"]
	groupDir := "deployments"
	mod, extractError := configs.ExtractModule(m, groupDir, configs.BACK)
	if extractError != nil {

		return extractError
	}
	path := "./" + groupDir + "/" + mod.Directory + "/helm"

	io.Println("path", path)
	arch := wrappers.ArchiveDir(path, m)

	// write archive to file
	io.Println("archive size", len(arch))

	wrappers.PushDir(arch)
	io.PrintColor("OK", io.Green)

	return nil
}
