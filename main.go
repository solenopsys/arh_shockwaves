package main

import (
	"github.com/spf13/cobra"
	"xs/commands/auth"
	"xs/commands/build"
	"xs/commands/chart"
	"xs/commands/cluster"
	"xs/commands/dev"
	"xs/commands/key"
	"xs/commands/net"
	"xs/commands/node"
	"xs/commands/public"
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
	rootCmd.AddCommand(dev.Cmd)
	rootCmd.AddCommand(auth.Cmd)
	rootCmd.AddCommand(build.Cmd)
}
