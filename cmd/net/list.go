package net

import (
	"github.com/spf13/cobra"
	"xs/internal/funcs"
	"xs/pkg/io"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List nodes of start network",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		records, err := funcs.GetSnRecords("solenopsys.org")
		if err != nil {
			io.Println(err)
		}
		for _, record := range records {
			io.Println(record)
		}
	},
}
