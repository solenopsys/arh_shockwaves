package services

import (
	"encoding/json"
	"os"
	"os/exec"
	"xs/pkg/io"
)

type FrontLib struct {
	cacheFile   string
	libs        map[string]string
	remoteCheck map[string]bool
}

func (b *FrontLib) genCache() {
	cmd := exec.Command("pnpm  ", []string{"cache"}...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		io.Println("Error executing command:", err)
	} else {
		io.Println("Output:", string(output))
	}
}

func (b *FrontLib) CacheCheck() {
	if _, err := os.Stat(b.cacheFile); os.IsNotExist(err) {
		b.genCache()
	}
}

func NewFrontLib() *FrontLib {
	return &FrontLib{libs: make(map[string]string), cacheFile: ".xs/cache.json"}
}

func tryDownLoadLib(fileName string) {

}

func tryUpLoadLib(fileName string) {

}

func (b *FrontLib) loadCache() {

	file, err := os.ReadFile(b.cacheFile)
	if err != nil {
		io.Panic(err)
	}

	err = json.Unmarshal(file, &b.libs)
	if err != nil {
		io.Panic(err)
	}

}
