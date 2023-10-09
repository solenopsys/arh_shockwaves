package tools

import (
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

func IpfsPublishDir(dir string, labels map[string]string) error {
	hosts := configs.GetInstanceConfManager().Conf.Hosts
	ipfsNode := wrappers.IpfsNode{IpfsNodeAddr: hosts.IpfsHost}
	cid, err := ipfsNode.UploadDirToIpfsNode(dir)
	pinning := wrappers.NewPinning()

	if err != nil {
		return err
	} else {
		io.Println("File cid: ", cid)
	}
	_, err = pinning.SmartPin(cid, labels)

	if err != nil {
		return err
	} else {
		io.Println("Pined cid: ", cid)
		return nil
	}
}
