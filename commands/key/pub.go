package key

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"xs/utils"
)

var cmdPubkey = &cobra.Command{
	Use:   "pub",
	Short: "Generate public key",
	Long:  `Print current time`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		seed := args[0]
		privateKey := utils.PrivateKeyFromSeed(seed)

		fmt.Println("Private key:", hex.EncodeToString(privateKey.Bytes()))
		fmt.Println("Public key:", hex.EncodeToString(privateKey.PubKey().Bytes()))

	},
}
