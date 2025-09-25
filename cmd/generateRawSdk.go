/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pulumi/provider-sdk-builder/pkg/builder"
)

// generateSdkCmd represents the generateSdk command
var generateSdkCmd = &cobra.Command{
	Use:   "generate-raw-sdk",
	Short: "Generates the source code for a Pulumi SDK",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Usage: 
generate-sdk --schema <schema file> --tfbridge <tfbridge binary file> --out <output directory> --language <language>`,
	Run: func(cmd *cobra.Command, args []string) {
		generateRawSdk(cmd, args)
	},
}

func generateRawSdk(cmd *cobra.Command, args []string) error {

	//TODO put in verbose flag?
	fmt.Printf("Generating SDK for provider %s\n SchemaPath: %s\n OutputPath: %s\n Languages: %v\n", providerName, schemaPath, outputPath, language)
	commands, err := builder.GenerateSdks(providerName, schemaPath, outputPath, language)
	if err != nil {
		return err
	}
	fmt.Printf("Prepared the following shell commands to run:\n\n%s\n\n", commands.String())
	output, err := builder.ExecuteCommandSequence(commands)
	fmt.Print(output)
	return err
}

func init() {
	generateSdkCmd.PersistentFlags().String("foo", "", "A help for foo")

	rootCmd.AddCommand(generateSdkCmd)
}
