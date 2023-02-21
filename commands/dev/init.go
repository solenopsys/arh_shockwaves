package dev

import (
	"github.com/spf13/cobra"
	"xs/services"
)

var cmdInit = &cobra.Command{
	Use:   "init [front/back]",
	Short: "Init monorepo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		services.LoadBase()
		services.NewHelper().InitRepository()
		getLoader(args).SyncFunc()
	},
}

func getLoader(args []string) *services.ConfLoader {
	if args[0] == "front" {
		return services.NewFrontLoader()
	} else if args[0] == "back" {
		return services.NewBackLoader()
	} else {
		panic("Unknown type needed (front/back)")
	}
}
