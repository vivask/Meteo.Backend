package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "xu4",
		Short: "xu4 is cmd app build with golang",
		Long:  `xu4 is golang xu4 service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
