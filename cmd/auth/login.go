package auth

import (
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
	"xs/internal/funcs"
	"xs/pkg/io"
	"xs/pkg/tools"
)

func readPassword() string {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		io.Println("Error reading password:", err)
		os.Exit(1)
	}
	io.Println()
	return string(bytePassword)
}

var cmdLogin = &cobra.Command{
	Use:   "login [username]",
	Short: "Authorisation",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		login := args[0]

		io.Println("Enter token:")
		password := readPassword()
		io.Println("Print password:", password)
		key := funcs.LoadKey(password, login)
		// hidden password read
		regData := funcs.UnMarshal(key)

		privateKey, err := tools.DecryptKeyData(regData.EncryptedKey, password)

		pk, err := tools.LoadPrivateKeyFromString(string(privateKey))
		if err != nil {
			io.Panic(err)
		}

		io.Println(pk)

		fileName, err := SOLENOPSYS_KEYS.WriteSessionToTempFile(key)

		if err != nil {
			io.Println("Error saving keys to file:", err)
			return
		}
		io.Println("Key saved to file", fileName)

	},
}
