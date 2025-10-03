/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pulumi/provider-sdk-builder/internal/builder"
)

// generateSdkCmd represents the generateSdk command
var generateSdkCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the source code for a Pulumi SDK",
	Long: `Generates sdks for each of the specified languages using the schema of your choice.
	Outputs will be stored in the form /sdk/{lang} in the directory specified by output path`,

	Run: func(cmd *cobra.Command, args []string) {
		generateRawSdk()
	},
}

var (
	generateOnlyInstructions = builder.BuildInstructions{GenerateSdks: true}
)

func generateRawSdk() error {

	//TODO refactor logging, put in verbose flag, clean up logs
	fmt.Printf("Generating SDK for provider %s\n SchemaPath: %s\n OutputPath: %s\n Languages: %v\n", providerName, schemaPath, outputPath, language)
	params := builder.BuildParameters{ProviderName: providerName, SchemaPath: schemaPath, OutputPath: outputPath, RawRequestedLanguage: language}
	commands, err := builder.GenerateBuildCmds(params, generateOnlyInstructions)
	if err != nil {
		return err
	}
	fmt.Printf("Prepared the following shell commands to run:\n\n%s\n\n", strings.Join(commands, "\n"))
	output, err := builder.ExecuteCommandSequence(commands)
	fmt.Print(output)
	return err
}

func init() {
	rootCmd.AddCommand(generateSdkCmd)
}
