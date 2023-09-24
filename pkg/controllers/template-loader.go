package controllers

import (
	"math/rand"
	"sync"
	"time"
	"xs/pkg/tools"
	"xs/pkg/wrappers/git"
)

type ModuleSourceLoader struct {
	servers []string
}

func NewSourceLoader() *ModuleSourceLoader {
	hostsNames := []string{"alpha", "bravo", "charlie"}
	servers := make([]string, len(hostsNames))
	nodesHost := "node.solenopsys.org"
	for i, hostName := range hostsNames {
		servers[i] = hostName + "." + nodesHost
	}
	rand.Seed(time.Now().UnixNano())
	return &ModuleSourceLoader{servers: servers}
}

func (t *ModuleSourceLoader) RandomServer() string {
	randomIndex := rand.Intn(len(t.servers))
	return t.servers[randomIndex]
}

func (t *ModuleSourceLoader) Load(cid string, path string, originalRemote string) error {
	url := "https://" + t.RandomServer() + "/ipns/" + cid

	wg := sync.WaitGroup{} // todo may be it not needed now
	wg.Add(1)
	err := git.CloneGitRepository(url, path, false, false, originalRemote)

	defer wg.Done()
	if err != nil {
		return err
	}
	gitDir := path + "/.git"
	return tools.DeleteDir(gitDir)
}
