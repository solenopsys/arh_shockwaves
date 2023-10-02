package publish

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdFile = &cobra.Command{
	Use:   "file [path] ",
	Short: "Publish file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		hosts := configs.GetInstanceConfManager().Conf.Hosts

		ipfsNode := wrappers.IpfsNode{IpfsNodeAddr: hosts.IpfsHost}

		cid, err := ipfsNode.UploadFileToIpfsNode(file)
		pinning := wrappers.NewPinning()
		labels := make(map[string]string)
		labels["type"] = "file"
		if err != nil {
			io.Println(err)
		} else {
			io.Println("File cid: ", cid)
		}
		_, err = pinning.SmartPin(cid, labels)

		if err != nil {
			io.Println(err)
		} else {
			io.Println("Pined!")
		}

	},
}
