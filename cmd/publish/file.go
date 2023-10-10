package publish

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdFile = &cobra.Command{
	Use:   "file [path] ",
	Short: "Publish file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		ipfsNode := wrappers.IpfsNode{IpfsNodeAddr: viper.GetString("hosts.ipfsNode")}

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
