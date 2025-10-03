/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

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

	//TODO put in verbose flag?
	fmt.Printf("Compiling SDK for provider %s\n OutputPath: %s\n Languages: %v\n", providerName, outputPath, language)
	params := builder.BuildParameters{OutputPath: outputPath, RawRequestedLanguage: language}

	// TODO don't actually do the work if we dont have the stuff we need to do it
	commands, err := builder.GenerateBuildCmds(params, compileOnlyInstructions)

	if err != nil {
		return err
	}
	fmt.Printf("Prepared the following shell commands to run:\n\n%s\n\n", strings.Join(commands, "\n"))
	// TODO the execution should not be driven by the command functions, they should just dispatch and be done with
	output, err := builder.ExecuteCommandSequence(commands)
	fmt.Print(output)
	return err
}

func init() {
	rootCmd.AddCommand(compileSdkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileSdkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileSdkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
