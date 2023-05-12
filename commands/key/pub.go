package key

import (
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
		utils.PrivateKeyFromSeed(seed)

		//bytesPrivateKey:=privateKey.
		//fmt.Println("Private key:", hex.EncodeToString(bytesPrivateKey))
		//fmt.Println("Public key:", hex.EncodeToString(privateKey.PubKey().Bytes()))

	},
}
