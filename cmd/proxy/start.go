package main

import "github.com/spf13/cobra"

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start proxy",
		Long:  `start proxy, default port is 5000`,
		Run:   startProxy,
	}
	enablePprof bool
)

func startProxy(cmd *cobra.Command, agrs []string) {

}
