package compilers

import (
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type HelmCompileExecutor struct {
	PrintConsole bool
}

func (n HelmCompileExecutor) Compile(params map[string]string) error {
	m := params["name"]
	groupDir := "modules"
	mod, extractError := configs.ExtractModule(m, groupDir, "back")
	if extractError != nil {

		return extractError
	}
	path := "./" + groupDir + "/" + mod.Directory + "/install"

	io.Println("path", path)
	arch := wrappers.ArchiveDir(path, m)

	// write archive to file
	io.Println("archive size", len(arch))

	wrappers.PushDir(arch)
	return nil
}
