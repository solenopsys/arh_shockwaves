package publish

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdFile = &cobra.Command{
	Use:   "file [path] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		ipfs := false
		hosts := configs.GetInstanceConfManager().Conf.Hosts

		if ipfs {
			cid, err := wrappers.UploadFileToIpfsNode(hosts.IpfsHost, file)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = file
			outChain, err := wrappers.UploadFileToIpfsCluster(hosts.IpfsClusterHost, d)

			if err != nil {
				io.Println(err)
			} else {
				//await chain
				io.Println("await chain")
				for out := range outChain {
					io.Println(out)
				}
			}
		}

	},
}
