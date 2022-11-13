package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server is cmd app build with golang",
		Long:  `server is golang server service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
