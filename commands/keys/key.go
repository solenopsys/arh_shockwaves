package keys

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/spf13/cobra"
)

// todo заглушка для импорта
var cmdAccount = &cobra.Command{
	Use:   "key",
	Short: "Gen key",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keyring := keys.NewDeleteKeyReq("bla")
		fmt.Println("Private Key:", keyring)
	},
}
