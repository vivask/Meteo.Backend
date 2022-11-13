package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "messanger",
		Short: "messanger is cmd app build with golang",
		Long:  `messanger is golang messanger service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
