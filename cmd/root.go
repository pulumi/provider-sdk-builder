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
	providerPath      string
	providerName      string
	rawLanguageString string
	outputPath        string
	sdkVersionString  string
	schemaPath        string

	verbose bool
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
	rootCmd.PersistentFlags().StringVarP(&providerPath, "providerPath", "p", "./", "Path to the provider you want to build")
	rootCmd.PersistentFlags().StringVarP(&providerName, "providerName", "n", "", "Name of the provider (required)")
	rootCmd.PersistentFlags().StringVarP(&rawLanguageString, "language", "l", "all", "Comma seperated list of programming languages you wish to generate SDKs for")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "out", "o", "", "Where you would like to output generated SDKs if different than {provider}/sdk")
	rootCmd.PersistentFlags().StringVar(&schemaPath, "schema", "", "Absolute path of schema.json. Defaults to  '{provider}/provider/cmd/pulumi-resource-random/schema.json'")
	rootCmd.PersistentFlags().StringVar(&sdkVersionString, "version", "4.0.0-alpha.0+dev", "SDK Version5")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	// Mark providerName as required
	rootCmd.MarkPersistentFlagRequired("providerName")
}
