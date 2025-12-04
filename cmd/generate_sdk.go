/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pulumi/provider-sdk-builder/internal/builder"
)

// generateSdkCmd represents the generateSdk command
var generateSdkCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the source code for a Pulumi SDK",
	Long:  `Generate SDK source code for specified languages from provider schema. Output to {sdkLocation}/{lang}/`,

	Run: func(cmd *cobra.Command, args []string) {
		generateRawSdk()
	},
}

var (
	generateOnlyInstructions = builder.BuildInstructions{GenerateSdks: true}
)

func generateRawSdk() error {

	if !quiet {
		fmt.Printf("Generating the SDKs for provider found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, providerName, rawLanguageString, schemaPath, sdkLocation, sdkVersionString)
	if err != nil {
		return err
	}

	return builder.BuildSDKs(params, generateOnlyInstructions, quiet, os.Stdout)
}

func init() {
	rootCmd.AddCommand(generateSdkCmd)
	generateSdkCmd.MarkFlagRequired("providerName")
}
