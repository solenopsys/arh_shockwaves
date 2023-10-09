package extractors

import (
	"strings"
)

type Microfrontend struct {
}

func (e Microfrontend) Extract(name string, path string) map[string]string {
	lib := strings.Replace(path, "./modules/", "@", 1)
	distribution := strings.Replace(path, "./", "./dist/", 1)

	params := map[string]string{ // todo remove this
		"lib":  lib,
		"dist": distribution,
		"name": name,
	}
	return params
}
