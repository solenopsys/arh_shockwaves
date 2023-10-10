package tools

import (
	"github.com/spf13/viper"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

func IpfsPublishDir(dir string, labels map[string]string) error {

	ipfsNode := wrappers.IpfsNode{IpfsNodeAddr: viper.GetString("hosts.ipfsNode")}
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
