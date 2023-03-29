package auth

import (
	"github.com/spf13/cobra"
	"log"
	"xs/utils"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Status of auth",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		jwtSessionBytes, err := JWT_SESSIONS.ReadSessionFromTempFile()
		if err != nil {
			log.Panic(err)
		}

		if jwtSessionBytes == nil {
			println("User not auth")
			return
		}

		keySaved, err := SOLENOPSYS_KEYS.ReadSessionFromTempFile()
		if err != nil {
			log.Panic(err)
		}

		regData := utils.UnMarshal(keySaved)

		pk, err := utils.LoadPublicKeyFromString(regData.PublicKey)
		if err != nil {
			log.Panic(err)
		}
		println(pk)
		//utils.JwtVerify(string(jwtSessionBytes), pk)
		println("SESSION", string(jwtSessionBytes))
	},
}
