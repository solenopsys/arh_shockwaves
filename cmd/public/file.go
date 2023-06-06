package public

import (
	"fmt"
	"github.com/spf13/cobra"
	"xs/pkg/wrappers"
)

var cmdFile = &cobra.Command{
	Use:   "file [path] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		ipfs := false

		if ipfs {
			cid, err := wrappers.UploadFileToIpfsNode("0.0.0.0:5003", file)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("File cid: ", cid)
			}
		} else {
			d := make([]string, 1)
			d[0] = file
			outChain, err := wrappers.UploadFileToIpfsCluster("0.0.0.0:9094", d)

			if err != nil {
				fmt.Println(err)
			} else {
				//await chain
				println("await chain")
				for out := range outChain {
					fmt.Println(out)
				}
			}
		}

	},
}
