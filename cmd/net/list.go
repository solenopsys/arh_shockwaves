package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"xs/internal/funcs"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List nodes of start network",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		records, err := funcs.GetSnRecords("solenopsys.org")
		if err != nil {
			fmt.Println(err)
		}
		for _, record := range records {
			fmt.Println(record)
		}
	},
}
