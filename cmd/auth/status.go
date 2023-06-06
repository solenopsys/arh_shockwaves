package auth

import (
	"github.com/spf13/cobra"
	"xs/internal/funcs"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of auth",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		jwtSessionBytes, err := JWT_SESSIONS.ReadSessionFromTempFile()
		if err != nil {
			io.Panic(err)
		}

		if jwtSessionBytes == nil {
			io.Println("User not auth")
			return
		}

		keySaved, err := SOLENOPSYS_KEYS.ReadSessionFromTempFile()
		if err != nil {
			io.Panic(err)
		}

		regData := funcs.UnMarshal(keySaved)

		pk, err := tools.LoadPublicKeyFromString(regData.PublicKey)
		if err != nil {
			io.Panic(err)
		}
		io.Println(pk)
		//utils.JwtVerify(string(jwtSessionBytes), pk)
		io.Println("SESSION", string(jwtSessionBytes))
	},
}
