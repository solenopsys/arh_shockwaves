package main

import (
	"github.com/spf13/cobra"
	"solenopsys-cli-xs/commands/cluster"
	"solenopsys-cli-xs/commands/keys"
	"solenopsys-cli-xs/commands/node"
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
	rootCmd.AddCommand(keys.Cmd)
	rootCmd.AddCommand(cluster.Cmd)
}
