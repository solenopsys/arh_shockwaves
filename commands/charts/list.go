package charts

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"solenopsys-cli-xs/utils"
)

var cmdList = cobra.Command{
	Use:   "list",
	Short: "List charts",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		api := utils.NewAPI(config)
		charts, err := api.ListHelmCharts("")

		if err != nil {
			log.Fatal(err)
		}
		for i, item := range charts.Items {
			fmt.Println(i, item.Name)
		}
		fmt.Println("Removed: ", args[0])
	},
}
