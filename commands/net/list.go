package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/utils"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List nodes of start network",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		records, err := utils.GetSnRecords("solenopsys.org")
		if err != nil {
			fmt.Println(err)
		}
		for _, record := range records {
			fmt.Println(record)
		}
	},
}
