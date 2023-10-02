package services

import (
	"encoding/json"
	"os"
	"os/exec"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type FrontLib struct {
	cacheFile       string
	npmCacheDir     string
	ipfsNode        *wrappers.IpfsNode
	libs            map[string]string
	remoteCheck     map[string]bool
	pinningRequests PinningRequests
	pinningService  *wrappers.Pinning
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
	return &FrontLib{libs: make(map[string]string), cacheFile: ".xs/cache.json", npmCacheDir: "node_modules/.cache/native-federation"}
}

func (b *FrontLib) filePath(fileName string) string {
	return b.npmCacheDir + "/" + fileName
}

func (b *FrontLib) tryDownLoadLib(fileName string) bool {
	cid, err := b.pinningRequests.FindFontLib(fileName)
	if err == nil {
		requests := NewIpfsRequests()
		fileBytes, err := requests.LoadCid(cid)
		if err != nil {
			io.Panic(err)
		}

		err = os.WriteFile(b.filePath(fileName), fileBytes, 444)
		if err != nil {
			io.Panic(err)
		}
		return true
	} else {
		if err != nil {
			io.Panic(err)
		}
	}
	return false
}

func (b *FrontLib) tryUpLoadLib(fileName string) (string, error) {
	cid, err := b.ipfsNode.UploadFileToIpfsNode(b.filePath(fileName))
	if err != nil {
		io.Panic(err)
	}
	labels := make(map[string]string)
	labels["front.static.library"] = fileName //todo const
	return b.pinningService.SmartPin(cid, labels)
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
