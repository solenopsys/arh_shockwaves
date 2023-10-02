package services

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"xs/pkg/io"
)

type IpfsRequests struct {
	servers []string
}

func NewIpfsRequests() *IpfsRequests {
	hostsNames := []string{"alpha", "bravo", "charlie"}
	servers := make([]string, len(hostsNames))
	nodesHost := "node.solenopsys.org"
	for i, hostName := range hostsNames {
		servers[i] = hostName + "." + nodesHost
	}
	rand.Seed(time.Now().UnixNano())
	return &IpfsRequests{servers: servers}
}

func (i *IpfsRequests) RandomServer() string {
	count := len(i.servers)
	randomIndex := rand.Intn(count)
	return i.servers[randomIndex]
}

func (i *IpfsRequests) GetCidUrl(cid string) string {
	return "https://" + i.RandomServer() + "/ipns/" + cid
}

func (i *IpfsRequests) LoadCid(cid string) ([]byte, error) {
	response, err := http.Get(i.GetCidUrl(cid))
	if err != nil {
		io.Println("Error:", err)
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)

}
