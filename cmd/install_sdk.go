/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pulumi/provider-sdk-builder/internal/builder"
	"github.com/spf13/cobra"
)

// installSdkCmd represents the installSdk command
var installSdkCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the compiled Pulumi SDKs locally for testing",
	Long: `Installs SDKs to the local development environment for testing purposes.

For each language, this command runs the appropriate installation steps:
  - nodejs: Creates a yarn link to the compiled SDK
  - dotnet: Copies .nupkg files to a local NuGet source and registers it
  - python: No installation steps (use pip install -e for local development)
  - go: No installation steps (Go uses local file paths directly)
  - java: No installation steps (use local Maven repository for testing)`,
	Run: func(cmd *cobra.Command, args []string) {
		installSdk()
	},
}

var (
	installOnlyInstructions = builder.BuildInstructions{InstallSdks: true}
)

func installSdk() error {

	if verbose {
		fmt.Printf("Installing the SDKs found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, rawLanguageString, schemaPath, outputPath, sdkVersionString)
	if err != nil {
		return err
	}

	commands, err := builder.GenerateBuildCmds(params, installOnlyInstructions)
	if err != nil {
		return err
	}

	output, err := builder.ExecuteCommandSequence(commands, verbose)
	fmt.Print(output)
	return err
}

func init() {
	rootCmd.AddCommand(installSdkCmd)
}
