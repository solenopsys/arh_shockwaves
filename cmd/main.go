package main

import (
	"github.com/spf13/cobra"
	"log"
	"xs/cmd/auth"
	"xs/cmd/build"
	"xs/cmd/chart"
	"xs/cmd/cluster"
	"xs/cmd/code"
	"xs/cmd/env"
	"xs/cmd/net"
	"xs/cmd/node"
	"xs/cmd/publish"
	"xs/cmd/serve"
	"xs/internal/configs"
	"xs/pkg/io"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	configs.GetInstanceConfManager() // read conf from xs directory

	var rootCmd = &cobra.Command{Use: "xs"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	initCommands(rootCmd)
	err := rootCmd.Execute()
	if err != nil {
		io.Fatal(err)
	}
}

func initCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(node.Cmd)
	rootCmd.AddCommand(cluster.Cmd)
	rootCmd.AddCommand(chart.Cmd)
	rootCmd.AddCommand(publish.Cmd)
	rootCmd.AddCommand(net.Cmd)
	rootCmd.AddCommand(env.Cmd)
	rootCmd.AddCommand(auth.Cmd)
	rootCmd.AddCommand(build.Cmd)
	rootCmd.AddCommand(code.Cmd)
	rootCmd.AddCommand(serve.Cmd)
}
