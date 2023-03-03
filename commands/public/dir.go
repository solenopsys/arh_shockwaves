package public

import (
	"fmt"
	"github.com/spf13/cobra"
	"xs/utils"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path]",
	Short: "Public dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		ipfs := false

		if ipfs {
			cid, err := utils.UploadDirToIpfsNode("0.0.0.0:5003", dir)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = dir
			cid, err := utils.UploadFileToIpfsCluster("0.0.0.0:9094", d)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("File cid: ", cid)
			}
		}

	},
}
