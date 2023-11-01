package services

import (
	"encoding/json"
	"os"
	"os/exec"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type FrontLibsController struct {
	cacheFile       string
	npmCacheDir     string
	ipfsNode        *wrappers.IpfsNode
	libs            map[string]string
	remoteCheck     map[string]bool
	pinningRequests *PinningRequests
	pinningService  *wrappers.Pinning
	key             string
}

const NPM_APPLICATION = "pnpm"

func (b *FrontLibsController) genCache() {
	cmd := exec.Command(NPM_APPLICATION, []string{"cache-synchronization"}...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		io.Println("Error executing command:", err)
	} else {
		store := io.GetLogStore()
		message := string(output)
		store.MessagesStream <- io.LogMessage{Message: message, Key: b.key}
		//println(message)
	}
}

func (b *FrontLibsController) CacheCheck() {
	if _, err := os.Stat(b.cacheFile); os.IsNotExist(err) {
		b.genCache()
	}
}

func NewFrontLibController(key string) *FrontLibsController {
	return &FrontLibsController{
		libs:            make(map[string]string),
		remoteCheck:     make(map[string]bool),
		cacheFile:       ".xs/to-upload.json",
		npmCacheDir:     "dist/shared",
		ipfsNode:        wrappers.NewIpfsNode(),
		pinningRequests: NewPinningRequests(),
		pinningService:  wrappers.NewPinning(),
		key:             key,
	}
}

func (b *FrontLibsController) filePath(fileName string) string {
	return b.npmCacheDir + "/" + fileName
}

func (b *FrontLibsController) tryDownLoadLib(fileName string) bool {
	//io.Println("Try download static front lib:", fileName)
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
	}
	return false
}

func (b *FrontLibsController) tryUpLoadLib(fileName string) (string, error) {
	io.Println("Upload static front lib:", fileName)

	cid, err := b.ipfsNode.UploadFileToIpfsNode(b.filePath(fileName))
	if err != nil {
		io.Panic(err)
	}
	labels := make(map[string]string)
	labels["front.static.library"] = fileName // todo move to const
	return b.pinningService.SmartPin(cid, labels)
}

func (b *FrontLibsController) loadCache() {

	file, err := os.ReadFile(b.cacheFile)
	if err != nil {
		io.Panic(err)
	}

	err = json.Unmarshal(file, &b.libs)
	if err != nil {
		io.Panic(err)
	}

}

func (b *FrontLibsController) localLibExists(fileName string) bool {
	path := b.filePath(fileName)
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func (b *FrontLibsController) PreProcessing() {

}

func (b *FrontLibsController) checkRemoteLib(fileName string) bool {
	if val, exists := b.remoteCheck[fileName]; exists {
		return val
	} else {
		_, err := b.pinningRequests.FindFontLib(fileName)
		exists := err == nil
		b.remoteCheck[fileName] = exists
		return exists
	}
}

func (b *FrontLibsController) PostProcessing() {
	for libName, fileName := range b.libs {
		libInLocalCache := b.localLibExists(fileName)
		libExistsInRemote := b.checkRemoteLib(fileName)
		if libInLocalCache && !libExistsInRemote {
			cid, err := b.tryUpLoadLib(fileName)
			io.Println("Upload lib:", libName, "file name:", fileName, "cid:", cid)
			if err != nil {
				io.Panic(err)
			}
			b.remoteCheck[fileName] = true
		}
	}
}
