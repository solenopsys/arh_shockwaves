package wrappers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	files "github.com/ipfs/boxo/files"
	ipfs "github.com/ipfs/go-ipfs-api"
	"io"
	"os"
	"path/filepath"
)

type IpfsNode struct {
	IpfsNodeAddr string
}

func (i *IpfsNode) UploadFileToIpfsNode(file string) (string, error) {
	sh := ipfs.NewShell(i.IpfsNodeAddr)

	fileBytes, err := os.ReadFile(file)
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

//func (i *IpfsNode) UploadFileToIpfsCluster(files []string) (chan api.AddedOutput, error) {
//	//split nodeAddr
//	split := strings.Split(i.IpfsNodeAddr, ":")
//
//	config := client.Config{
//		Host: split[0],
//		Port: split[1],
//	}
//	clusterClient, err := client.NewDefaultClient(&config)
//	if err != nil {
//		lio.Panic(err)
//	}
//
//	outChain := make(chan api.AddedOutput, 1)
//	// Add the files to IPFS Cluster
//	err = clusterClient.Add(context.Background(), files, api.DefaultAddParams(), outChain)
//	if err != nil {
//		lio.Panic(err)
//	}
//
//	return outChain, nil
//}

func Hidden(enabled bool) ipfs.AddOpts {
	return func(rb *ipfs.RequestBuilder) error {
		rb.Option("hidden", enabled)
		return nil
	}
}

func addDirRecursiveIncludeHidden(s *ipfs.Shell, dir string, options ...ipfs.AddOpts) (string, error) { // todo need push to ipfs
	stat, err := os.Lstat(dir)
	if err != nil {
		return "", err
	}

	sf, err := files.NewSerialFile(dir, true, stat)
	if err != nil {
		return "", err
	}
	slf := files.NewSliceDirectory([]files.DirEntry{files.FileEntry(filepath.Base(dir), sf)})
	reader := files.NewMultiFileReader(slf, true, true)

	rb := s.Request("add").Option("recursive", true)
	for _, option := range options {
		option(rb)
	}

	// Here we cannot use .Exec because "add" streams responses back for each file
	// within the directory, and we only care about the last one, which is the directory
	// itself.
	resp, err := rb.Body(reader).Send(context.Background())
	if err != nil {
		return "", err
	}
	defer resp.Close()

	if resp.Error != nil {
		return "", resp.Error
	}

	dec := json.NewDecoder(resp.Output)
	var final string
	for {
		var out object
		err = dec.Decode(&out)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		final = out.Hash
	}

	if final == "" {
		return "", errors.New("no results received")
	}

	return final, nil

}

type object struct {
	Hash string
}

func (i *IpfsNode) UploadDirToIpfsNode(dir string) (string, error) {

	sh := ipfs.NewShell(i.IpfsNodeAddr)

	//cid, err := sh.AddDir(dir)
	cid, err := addDirRecursiveIncludeHidden(sh, dir)

	if err != nil {
		return "", err
	} else {
		return cid, nil
	}
}
