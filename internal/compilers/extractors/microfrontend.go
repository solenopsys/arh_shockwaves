package extractors

import (
	"strings"
)

type Microfrontend struct {
}

func (e Microfrontend) Extract(name string, path string) map[string]string {
	path = strings.Replace(path, "./modules/", "@", 1)
	distribution := strings.Replace(path, "./frontends/", "./frontends/dist/", 1)

	params := map[string]string{ // todo remove this
		"path": path,
		"dist": distribution,
		"name": name,
	}
	return params
}
