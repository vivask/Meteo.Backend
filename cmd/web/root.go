package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server is web cmd app build with golang",
		Long:  `server is web golang web service using gin and gorm`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
