package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "upssched-cmd",
		Short: "upssched-cmd is cmd app build with golang",
		Long:  `upssched-cmd is golang upssched-cmd command`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
