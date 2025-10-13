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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		compileSdk()
	},
}

var (
	compileOnlyInstructions = builder.BuildInstructions{CompileSdks: true}
)

func compileSdk() error {

	if verbose {
		fmt.Printf("Compiling the SDKs found at Path: %s\nLanguages: %v\n", providerPath, rawLanguageString)
	}

	params, err := builder.ParseInputs(providerPath, providerName, rawLanguageString, schemaPath, outputPath, sdkVersionString)
	if err != nil {
		return err
	}

	commands, err := builder.GenerateBuildCmds(params, compileOnlyInstructions)
	if err != nil {
		return err
	}

	return builder.ExecuteCommandSequence(commands, verbose, os.Stdout)
}

func init() {
	rootCmd.AddCommand(compileSdkCmd)
}
