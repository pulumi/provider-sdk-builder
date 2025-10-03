package lang

import (
	"strings"
)

const (
	joinCmdLineEnding             = " && \\\n"
	pulumiPackageGenSdkCmd string = "pulumi package gen-sdk {Path} --language {Lang}"
	outputDirectoryString  string = " --out {OutputPath}"
)

// BaseGenerateSdkCommand creates a shell command for generating an SDK
// using the baseGenerateCmd template with the provided parameters
func BaseGenerateSdkCommand(providerName, path, outputPath, langString string) []string {
	cmd := pulumiPackageGenSdkCmd
	cmd = strings.ReplaceAll(cmd, "{ProviderName}", providerName)
	cmd = strings.ReplaceAll(cmd, "{Path}", path)
	cmd = strings.ReplaceAll(cmd, "{Lang}", langString)
	if outputPath != "" {
		outputDir := outputDirectoryString
		outputDir = strings.ReplaceAll(outputDir, "{OutputPath}", outputPath)
		cmd += outputDir
	}

	return []string{cmd}
}
