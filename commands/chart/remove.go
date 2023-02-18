package chart

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"xs/utils"
)

var cmdRemove = &cobra.Command{
	Use:   "remove [chart]",
	Short: "Module chart",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		api := utils.NewAPI(config)
		api.DeleteHelmChart(args[0])

		fmt.Println("Removed: ", args[0])
	},
}
