package main

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var cmdKeys = &cobra.Command{
	Use:   "keys [command]",
	Short: "Keys manipulation functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

var cmdNode = &cobra.Command{
	Use:   "node [command]",
	Short: "Node control functions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + args[0])
	},
}

var cmdNodeInstall = &cobra.Command{
	Use:   "install [connect to node]",
	Short: "Install node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start install")

		httpBody, err := downloadScript()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		cmdIn := exec.Command("sh")
		cmdIn.Stdin = bytes.NewBuffer(httpBody)

		stdout, err := cmdIn.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}

		// Start the command
		err = cmdIn.Start()
		if err != nil {
			fmt.Println(err)
		}

		// Use io.Copy to print the command's output in real-time
		_, err = io.Copy(os.Stdout, stdout)
		if err != nil {
			fmt.Println(err)
		}

		// Wait for the command to finish
		err = cmdIn.Wait()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func downloadScript() ([]byte, error) {
	response, err := http.Get("https://get.k3s.io")
	if err != nil {
		return nil, err
	} else {
		return ioutil.ReadAll(response.Body)
	}
}

var cmdNodeRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start removing ")
	},
}

var cmdNodeStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of node",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start removing ")
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

func main() {

	cmdKeys.AddCommand(cmdPubkey)
	cmdKeys.AddCommand(cmdSeed)

	cmdNode.AddCommand(cmdNodeInstall)
	var rootCmd = &cobra.Command{Use: "xs"}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.AddCommand(cmdKeys)
	rootCmd.AddCommand(cmdNode)

	rootCmd.Execute()
}
