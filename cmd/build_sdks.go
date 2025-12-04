/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/pulumi/provider-sdk-builder/internal/builder"
	"github.com/spf13/cobra"
)

// buildSdksCmd represents the build-sdks command
var buildSdksCmd = &cobra.Command{
	Use:   "build-sdks",
	Short: "Generates, compiles, and packages SDKs for use",
	Long:  `Generates and compiles SDKs from schema`,
	Run: func(cmd *cobra.Command, args []string) {
		buildSdk()
	},
}

var (
	buildAllInstructions = builder.BuildInstructions{GenerateSdks: true, CompileSdks: true}
)

func buildSdk() error {

	if !quiet {
		fmt.Printf("Building SDKs for provider found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, providerName, rawLanguageString, schemaPath, sdkLocation, sdkVersionString)
	if err != nil {
		return err
	}

	return builder.BuildSDKs(params, buildAllInstructions, quiet, os.Stdout)
}

func init() {
	rootCmd.AddCommand(buildSdksCmd)
	buildSdksCmd.MarkFlagRequired("providerName")
}
