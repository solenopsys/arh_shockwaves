package public

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdName = &cobra.Command{
	Use:   "name [name] [cid]",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		cid := args[1]

		pinning := &wrappers.Pinning{}

		pinning.Host = "http://" + PinningHost
		pinning.UserKey = "alexstorm" // todo remove it

		labels := make(map[string]string)
		labels["code.site"] = name

		err := pinning.SmartPin(cid, labels, name)

		if err != nil {
			io.Println(err)
		}

	},
}
