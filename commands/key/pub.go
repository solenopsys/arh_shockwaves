package key

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
)

var cmdPubkey = &cobra.Command{
	Use:   "pub",
	Short: "Generate public key",
	Long:  `Print current time`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		seed := []byte(args[0])
		masterKey, _ := bip32.NewMasterKey(seed)
		publicKey := masterKey.PublicKey()
		//fmt.Println("Private key:", masterKey.String())
		fmt.Println("Public key\n", publicKey.String())

	},
}
