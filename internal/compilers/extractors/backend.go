package extractors

import (
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Backend struct {
}

func (e Backend) Extract(name string, path string) map[string]string {

	res, rootDir := tools.FindWorkspaceRootDir()
	if !res {
		io.Fatal("Workspace root dir not found")
	}

	params := map[string]string{
		"name": name,
		"path": path,
		"dist": rootDir + "/shockwaves/dist",
	}
	return params
}
