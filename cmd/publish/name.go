package publish

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"xs/internal/configs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdName = &cobra.Command{
	Use:   "name [name] [cid]",
	Short: "Publish inpn name for cid",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		cid := args[1]

		pinning := &wrappers.Pinning{}

		pinning.Host = viper.GetString("hosts.pinningService")
		pinning.UserKey = configs.GetAuthManager().PublicKey

		labels := make(map[string]string)
		labels["code.site"] = name

		err := pinning.SmartPinAndName(cid, labels, name)

		if err != nil {
			io.Println(err)
		}

	},
}
