package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sshclient",
		Short: "sshclient is cmd app build with golang",
		Long:  `sshclient is golang sshclient service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
