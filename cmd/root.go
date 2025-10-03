/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "provider-sdk-builder",
	Short: "Generate Pulumi SDKs for a given cloud provider",
	Long: `Generate Pulumi Provider SDKs against a given cloud provider
Supports Go, Python, .NET, NodesJS, and Java. 
SDKs require a schema file. Terraform backed providers also require a tfbridge binary.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	providerName     string
	schemaPath       string
	language         string
	outputPath       string
	sdkVersionString string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Global flags used in multiple commands
	rootCmd.PersistentFlags().StringVarP(&providerName, "providerName", "p", "", "Name of the provider, e.g. 'aws'")
	rootCmd.PersistentFlags().StringVarP(&schemaPath, "schema", "s", "", "Path to the directory with the schema file and tfbridge binary (if applicable)")
	rootCmd.PersistentFlags().StringVarP(&language, "language", "l", "all", "Programming language")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "out", "o", "", "Output directory")
	rootCmd.PersistentFlags().StringVarP(&sdkVersionString, "version", "v", "4.0.0-alpha.0+dev", "SDK Version")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
