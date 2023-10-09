package services

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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

func (i *IpfsRequests) GetNameCidUrl(cid string) string {
	return "https://" + i.RandomServer() + "/ipns/" + cid
}

func (i *IpfsRequests) GetCidUrl(cid string) string {
	return "https://" + i.RandomServer() + "/ipfs/" + cid
}

func (i *IpfsRequests) LoadCid(cid string) ([]byte, error) {
	response, err := http.Get(i.GetCidUrl(cid))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("ipfs node return not 200 status code")
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)

}
