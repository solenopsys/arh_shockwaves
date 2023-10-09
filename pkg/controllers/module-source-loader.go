package controllers

import (
	"sync"
	"xs/internal/services"
	"xs/pkg/tools"
	"xs/pkg/wrappers/git"
)

type ModuleSourceLoader struct {
	ipfs *services.IpfsRequests
}

func (t *ModuleSourceLoader) Load(cid string, path string, originalRemote string) error {

	wg := sync.WaitGroup{} // todo may be it not needed now
	wg.Add(1)
	err := git.CloneGitRepository(t.ipfs.GetNameCidUrl(cid), path, false, false, false, originalRemote)

	defer wg.Done()
	if err != nil {
		return err
	}
	gitDir := path + "/.git"
	if originalRemote == "" {
		return tools.DeleteDir(gitDir)
	} else {
		return nil
	}
}

func NewModuleSourceLoader() *ModuleSourceLoader {
	return &ModuleSourceLoader{ipfs: services.NewIpfsRequests()}
}
