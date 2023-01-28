package key

import (
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
)

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
