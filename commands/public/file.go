package public

import (
	"bytes"
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"xs/utils"
)

var cmdFile = &cobra.Command{
	Use:   "file [path] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		sh := ipfs.NewShell("0.0.0.0:5003")

		fileBytes, err := utils.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add the file to IPFS
		cid, err := sh.Add(bytes.NewReader(fileBytes))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File cid: ", cid)
	},
}
