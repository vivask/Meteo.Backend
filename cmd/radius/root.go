package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "radius",
		Short: "radius is cmd app build with golang",
		Long:  `radius is golang radius service`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
