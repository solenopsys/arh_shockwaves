package configs

import (
	"encoding/json"
	"xs/pkg/tools"
)

type NpmLibPackage struct {
	Name                       string            `json:"name"`
	AllowedNonPeerDependencies map[string]string `json:"allowedNonPeerDependencies"`
	Dependencies               map[string]string `json:"dependencies"`
	PeerDependencies           map[string]string `json:"peerDependencies"`
}

func LoadNpmLibPackage(packageJsonPath string) (lp *NpmLibPackage) {
	lp = &NpmLibPackage{}
	bytesFromFile, err := tools.ReadFile(packageJsonPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), lp)
	if err != nil {
		panic(err)
	}
	return lp
}