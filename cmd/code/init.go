package code

import (
	"github.com/spf13/cobra"
	"log"
	"xs/cmd/public"
	"xs/pkg/wrappers"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Workspace initialization",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pinning := wrappers.Pinning{}
		pinning.Host = "http://" + public.PinningHost // todo remove it
		pinning.UserKey = "alexstorm"                 // todo remove it

		repo, err := pinning.FindRepo("@solenopsys/tp-workspace")

		if err != nil {
			log.Fatal(err)
		}

		println(repo)

		//	configs.LoadWorkspace("https://github.com/solenopsys/tp-workspace.git")
	},
}
