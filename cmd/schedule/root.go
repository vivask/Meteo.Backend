package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "schedule",
		Short: "schedule is cmd app build with golang",
		Long:  `schedule is golang schedule service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
