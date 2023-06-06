package chart

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"xs/pkg/wrappers"
)

var cmdInstall = &cobra.Command{
	Use:   "install [chart] [version] [repository]",
	Short: "Install chart",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		kuber := wrappers.Kuber{}

		config, err := kuber.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		api := wrappers.NewAPI(config)
		chart := args[0]
		repoUrl := args[2]
		version := args[1]
		simple, err := api.CreateHelmChartSimple(chart, repoUrl, version, "default")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Installed: ", simple.Name)
	},
}
