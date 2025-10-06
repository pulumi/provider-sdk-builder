/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildSdksCmd represents the build-sdks command
var buildSdksCmd = &cobra.Command{
	Use:   "build-sdks",
	Short: "Generates, compiles, and packages SDKs for use",
	Long:  `Accepts a schema file, tfbridge binary (if applicable), output directory, and language to generate, compile, and package the SDK for use.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildSdks called")
	},
}

func init() {
	rootCmd.AddCommand(buildSdksCmd)
}
