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
Supports Go, Python, .NET, NodeJS, and Java.
SDKs require a schema file. Terraform backed providers also require a tfbridge binary.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	providerPath      string
	providerName      string
	rawLanguageString string
	sdkLocation       string
	sdkVersionString  string
	schemaPath        string

	quiet bool
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
	rootCmd.PersistentFlags().StringVarP(&providerPath, "providerPath", "p", "./", "Path to provider directory")
	rootCmd.PersistentFlags().StringVarP(&providerName, "providerName", "n", "", "Name of the provider (required)")
	rootCmd.PersistentFlags().StringVarP(&rawLanguageString, "language", "l", "all", "Languages to generate (comma-separated or 'all')")
	rootCmd.PersistentFlags().StringVar(&sdkLocation, "sdkLocation", "", "SDK directory (default: {provider}/sdk)")
	rootCmd.PersistentFlags().StringVar(&schemaPath, "schema", "", "Path to schema.json file")
	rootCmd.PersistentFlags().StringVar(&sdkVersionString, "version", "4.0.0-alpha.0+dev", "SDK Version")

	// Deprecated alias for backwards compatibility
	rootCmd.PersistentFlags().StringVarP(&sdkLocation, "out", "o", "", "Deprecated: use --sdkLocation instead")
	rootCmd.PersistentFlags().MarkDeprecated("out", "please use --sdkLocation instead")

	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress command output")
}
