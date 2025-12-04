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

// compileSdkCmd represents the compileSdk command
var compileSdkCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile and package Pulumi SDKs",
	Long:  `Compile and package generated SDKs for distribution`,
	Run: func(cmd *cobra.Command, args []string) {
		compileSdk()
	},
}

var (
	compileOnlyInstructions = builder.BuildInstructions{CompileSdks: true}
)

func compileSdk() error {

	if !quiet {
		fmt.Printf("Compiling the SDKs found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, providerName, rawLanguageString, schemaPath, sdkLocation, sdkVersionString)
	if err != nil {
		return err
	}

	return builder.BuildSDKs(params, compileOnlyInstructions, quiet, os.Stdout)
}

func init() {
	rootCmd.AddCommand(compileSdkCmd)
	compileSdkCmd.MarkFlagRequired("providerName")
}
