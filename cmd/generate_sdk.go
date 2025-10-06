/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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

	if verbose {
		fmt.Printf("Generating the SDKs for provider found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, rawLanguageString, schemaPath, outputPath, sdkVersionString)
	if err != nil {
		return err
	}

	commands, err := builder.GenerateBuildCmds(params, generateOnlyInstructions)
	if err != nil {
		return err
	}

	output, err := builder.ExecuteCommandSequence(commands, verbose)
	fmt.Print(output)
	return err
}

func init() {
	rootCmd.AddCommand(generateSdkCmd)
}
