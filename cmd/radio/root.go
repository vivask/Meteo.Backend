package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "radio",
		Short: "radio is cmd app build with golang",
		Long:  `radio is golang messanger service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
