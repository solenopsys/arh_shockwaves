package public

import (
	"github.com/spf13/cobra"
	"xs/pkg/controllers"
	"xs/pkg/io"
)

var cmdGit = &cobra.Command{
	Use:   "sync [config] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]

		pg := &controllers.PublicGit{
			IpfsHost:    IpfsHost,
			PinningHost: PinningHost,
		}

		err := pg.LoadConfig(configName)

		if err != nil {
			io.Println(err)
			return
		}

		pg.ProcessingFile("solenopsys") // todo remove it

		//
	},
}
