package public

import (
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path] ",
	Short: "Public dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		sh := ipfs.NewShell("0.0.0.0:5003")

		cid, err := sh.AddDir(dir)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Directory cid: ", cid)
		}

	},
}
