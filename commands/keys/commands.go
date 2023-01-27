package keys

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
)

var Cmd = &cobra.Command{
	Use:   "keys [command]",
	Short: "Keys manipulation functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

var cmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Generate seed",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		fmt.Println("Seed phrase\n", mnemonic)

	},
}

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

// todo заглушка для импорта
var cmdAccount = &cobra.Command{
	Use:   "key",
	Short: "Gen key",
	Long:  `Print current time`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keyring := keys.NewDeleteKeyReq("bla")
		fmt.Println("Private Key:", keyring)
	},
}

func init() {
	Cmd.AddCommand(cmdPubkey)
	Cmd.AddCommand(cmdSeed)
	Cmd.AddCommand(cmdAccount)
}
