package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "healthy",
		Short: "healthy is cmd app build with golang",
		Long:  `healthy is golang healthy command`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
