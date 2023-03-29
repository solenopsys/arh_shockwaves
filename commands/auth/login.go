package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
	"xs/utils"
)

func readPassword() string {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}
	fmt.Println()
	return string(bytePassword)
}

var cmdLogin = &cobra.Command{
	Use:   "login [username]",
	Short: "Authorisation",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		login := args[0]

		println("Enter password:")
		password := readPassword()
		println("Print password:", password)
		key := utils.LoadKey(password, login)
		// hidden password read
		regData := utils.UnMarshal(key)

		privateKey, err := utils.DecryptKeyData(regData.EncryptedKey, password)

		pk, err := utils.LoadPrivateKeyFromString(string(privateKey))
		if err != nil {
			log.Panic(err)
		}

		println(pk)

		fileName, err := SOLENOPSYS_KEYS.WriteSessionToTempFile(key)

		if err != nil {
			println("Error saving keys to file:", err)
			return
		}
		println("Key saved to file", fileName)

	},
}
