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

var (
	outDir string
)

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildSdksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildSdksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildSdksCmd.Flags().StringVarP(&outDir, "out", "o", "./sdk", "Output directory")

	rootCmd.AddCommand(buildSdksCmd)
}
