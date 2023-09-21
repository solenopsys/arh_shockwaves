package public

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path]",
	Short: "Public dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		ipfs := true
		hosts := configs.GetInstanceConfManager().Conf.Hosts

		if ipfs {
			cid, err := wrappers.UploadDirToIpfsNode(hosts.IpfsHost, dir)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = dir
			cid, err := wrappers.UploadFileToIpfsCluster(hosts.IpfsClusterHost, d)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		}

	},
}
