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
	Short: "Invoke to build Pulumi Provider SDKs",
	Long: `
Pulumi Provider SDK Builder generates SDKs that allow users to build Pulumi programs in the language of their choice.
This tool currently supports generating Go, Python, .NET, NodesJS, and Java SDKs. 
Each SDK is generated from a schema file and a Terraform Bridge Binary file for Terraform backed providers.
Usage generateSDK --schema <schema file> --tfbridge <tfbridge binary file> --out <output directory> --language <language>
All arguments are optional, and if none are provided the tool will attempt to generate all SDKs, and will look for a schema and tfbridge in the local file structure.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.provider-sdk-builder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
