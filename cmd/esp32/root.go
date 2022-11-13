package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "esp32",
		Short: "esp32 is cmd app build with golang",
		Long:  `esp32 is golang esp32 service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
