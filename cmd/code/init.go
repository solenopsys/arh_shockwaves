package code

import (
	"github.com/spf13/cobra"
	"log"
	"xs/internal/configs"
	"xs/pkg/wrappers"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Workspace initialization",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pinning := wrappers.NewPinning()
		repo, err := pinning.FindOne("@solenopsys/tp-workspace")
		if err != nil {
			log.Fatal(err)
		}
		tl := configs.NewSourceLoader()
		tl.Load(repo.Cid, ".") // todo random from config
	},
}
