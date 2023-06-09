package wrappers

import (
	"bytes"
	"context"
	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
	ipfs "github.com/ipfs/go-ipfs-api"
	"strings"
	"xs/pkg/io"
	"xs/pkg/tools"
)

func UploadFileToIpfsNode(nodeAddr string, file string) (string, error) {
	sh := ipfs.NewShell(nodeAddr)

	fileBytes, err := tools.ReadFile(file)
	if err != nil {
		return "", err
	}

	// Add the file to IPFS
	cid, err := sh.Add(bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}
	return cid, nil
}

func UploadFileToIpfsCluster(nodeAddr string, files []string) (chan api.AddedOutput, error) {
	//split nodeAddr
	split := strings.Split(nodeAddr, ":")
	config := client.Config{
		Host: split[0],
		Port: split[1],
	}
	clusterClient, err := client.NewDefaultClient(&config)
	if err != nil {
		io.Panic(err)
	}

	outChain := make(chan api.AddedOutput, 1)
	// Add the files to IPFS Cluster
	err = clusterClient.Add(context.Background(), files, api.DefaultAddParams(), outChain)
	if err != nil {
		io.Panic(err)
	}

	return outChain, nil
}

func UploadDirToIpfsNode(nodeAddr string, dir string) (string, error) {

	sh := ipfs.NewShell(nodeAddr)

	cid, err := sh.AddDir(dir)
	if err != nil {
		return "", err
	} else {
		return cid, nil
	}
}
