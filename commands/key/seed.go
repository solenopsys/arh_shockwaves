package key

import (
	"fmt"
	"github.com/spf13/cobra"
	"xs/utils"
)

var cmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Generate seed",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		mnemonic := utils.GenMnemonic()
		fmt.Println("Seed phrase\n", mnemonic)

	},
}
