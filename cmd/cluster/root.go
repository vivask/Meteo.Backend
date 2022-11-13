package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cluster",
		Short: "cluster is cmd app build with golang",
		Long:  `cluster is golang cluster service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
