package main

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/commands/chart"
	"solenopsys-cli-xs/commands/cluster"
	"solenopsys-cli-xs/commands/key"
	"solenopsys-cli-xs/commands/net"
	"solenopsys-cli-xs/commands/node"
	"solenopsys-cli-xs/commands/public"
)

func main() {
	var rootCmd = &cobra.Command{Use: "xs"}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	initCommands(rootCmd)

	rootCmd.Execute()
}

func initCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(node.Cmd)
	rootCmd.AddCommand(key.Cmd)
	rootCmd.AddCommand(cluster.Cmd)
	rootCmd.AddCommand(chart.Cmd)
	rootCmd.AddCommand(public.Cmd)
	rootCmd.AddCommand(net.Cmd)
}
