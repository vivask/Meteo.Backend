package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "media",
		Short: "media is cmd app build with golang",
		Long:  `media is golang media service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
