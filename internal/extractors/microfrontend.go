package extractors

import (
	"strings"
)

type Microfrontend struct {
}

func (e Microfrontend) Extract(name string, path string) map[string]string {
	params := map[string]string{ // todo remove this
		"lib": strings.Replace(path, "./modules/", "@", 1),
		//	"dest": dest,
		"name": name,
	}
	return params
}
