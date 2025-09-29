/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageSdkCmd = &cobra.Command{
	Use:   "package",
	Short: "Cleans up unneeded source code and packages files for release",
	Long:  `Cleans up unneeded source code and packages files for release.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("package for release called")
	},
}

func init() {
	rootCmd.AddCommand(packageSdkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageForReleaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packageForReleaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
