package publish

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdDir = &cobra.Command{
	Use:   "dir [path]",
	Short: "Publish dir in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		labels := make(map[string]string)
		labels["type"] = "dir"

		cid, err := tools.IpfsPublishDir(dir, labels)

		if err != nil {
			io.Println(err)
		} else {
			io.Println("Pined! ", cid)
		}
	},
}
