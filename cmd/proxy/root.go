package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "proxy",
		Short: "proxy is dsn cmd app build with golang",
		Long:  `proxy is golang dns service with support ad block list and unlock https`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
