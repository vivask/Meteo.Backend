package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "nut",
		Short: "nut is cmd app build with golang",
		Long:  `nut is golang nut service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
