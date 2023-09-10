package configs

import (
	"math/rand"
	"sync"
	"time"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
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

func (t *ModuleSourceLoader) Load(cid string, path string) {
	url := "https://" + t.RandomServer() + "/ipns/" + cid

	wg := sync.WaitGroup{} // todo may be it not needed now
	wg.Add(1)
	err := wrappers.CloneGitRepository(url, path, false, false)

	defer wg.Done()
	if err != nil {
		io.Panic(err)
	} else {
		gitDir := path + "/.git"
		err := tools.DeleteDir(gitDir)
		if err != nil {
			io.Panic(err)
		}
	}
}
