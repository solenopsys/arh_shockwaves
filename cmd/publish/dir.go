package publish

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path]",
	Short: "Publish dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		ipfs := true
		hosts := configs.GetInstanceConfManager().Conf.Hosts

		ipfsNode := wrappers.IpfsNode{IpfsNodeAddr: hosts.IpfsHost}

		if ipfs {
			cid, err := ipfsNode.UploadDirToIpfsNode(dir)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = dir
			cid, err := ipfsNode.UploadFileToIpfsCluster(d)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		}

	},
}
