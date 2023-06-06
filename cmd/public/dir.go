package public

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path]",
	Short: "Public dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		ipfs := false

		if ipfs {
			cid, err := wrappers.UploadDirToIpfsNode("0.0.0.0:5003", dir)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = dir
			cid, err := wrappers.UploadFileToIpfsCluster("0.0.0.0:9094", d)

			if err != nil {
				io.Println(err)
			} else {
				io.Println("File cid: ", cid)
			}
		}

	},
}
