package extractors

import (
	"strings"
	"xs/internal/configs"
)

type Microfrontend struct {
}

func (e Microfrontend) Extract(name string, path string) map[string]string {
	path = strings.Replace(path, "./modules/", "@", 1)

	lp := configs.LoadNpmLibPackage(path + "/package.json")

	distribution := strings.Replace(path, "./converged/", "./converged/dist/", 1)

	params := map[string]string{ // todo remove this
		"path":    path,
		"dist":    distribution,
		"name":    name,
		"version": lp.Version,
	}
	return params
}
