package lang

import (
	"strings"
)

const (
	baseGenerateSdkCmd string = "{Path}/bin/pulumi-tfgen-{ProviderName} {Lang} --out {OutputPath}/{Lang}/"
)

// BaseGenerateSdkCommand creates a shell command for generating an SDK
// using the baseGenerateCmd template with the provided parameters
func BaseGenerateSdkCommand(providerName, path, outputPath, langString string) []string {
	cmd := baseGenerateSdkCmd
	cmd = strings.ReplaceAll(cmd, "{ProviderName}", providerName)
	cmd = strings.ReplaceAll(cmd, "{Path}", path)
	cmd = strings.ReplaceAll(cmd, "{OutputPath}", outputPath)
	cmd = strings.ReplaceAll(cmd, "{Lang}", langString)

	return []string{cmd}
}
