package configs

import (
	"encoding/json"
	"os"
	"xs/pkg/io"
)

type NpmLibPackage struct {
	Name                       string            `json:"name"`
	Version                    string            `json:"version"`
	AllowedNonPeerDependencies map[string]string `json:"allowedNonPeerDependencies"`
	Dependencies               map[string]string `json:"dependencies"`
	PeerDependencies           map[string]string `json:"peerDependencies"`
}

func LoadNpmLibPackage(packageJsonPath string) (lp *NpmLibPackage) {
	lp = &NpmLibPackage{}
	bytesFromFile, err := os.ReadFile(packageJsonPath)

	if err != nil {
		io.Panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), lp)
	if err != nil {
		io.Panic(err)
	}
	return lp
}
