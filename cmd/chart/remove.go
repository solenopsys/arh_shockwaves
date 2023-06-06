package chart

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"xs/pkg/wrappers"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [chart]",
	Short: "Module chart",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kuber := wrappers.Kuber{}

		config, err := kuber.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		api := wrappers.NewAPI(config)
		api.DeleteHelmChart(args[0])

		fmt.Println("Removed: ", args[0])
	},
}
