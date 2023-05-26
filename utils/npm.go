package utils

import "encoding/json"

type NpmLibPackage struct {
	Name                       string            `json:"name"`
	AllowedNonPeerDependencies map[string]string `json:"allowedNonPeerDependencies"`
}

func LoadNpmLibPackage(packageJson string) (lp *NpmLibPackage) {
	lp = &NpmLibPackage{}
	err := json.Unmarshal([]byte(packageJson), lp)
	if err != nil {
		panic(err)
	}
	return lp
}
