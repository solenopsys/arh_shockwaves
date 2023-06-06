package main

import (
	"github.com/spf13/cobra"
	"xs/cmd/auth"
	"xs/cmd/build"
	"xs/cmd/chart"
	"xs/cmd/cluster"
	"xs/cmd/env"
	"xs/cmd/net"
	"xs/cmd/node"
	"xs/cmd/public"
	"xs/cmd/ws"
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
	rootCmd.AddCommand(cluster.Cmd)
	rootCmd.AddCommand(chart.Cmd)
	rootCmd.AddCommand(public.Cmd)
	rootCmd.AddCommand(net.Cmd)
	rootCmd.AddCommand(env.Cmd)
	rootCmd.AddCommand(auth.Cmd)
	rootCmd.AddCommand(build.Cmd)
	rootCmd.AddCommand(ws.Cmd)
}
