package lang

import (
	"strings"

	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

const (
	baseGenerateSdkCmd string = "{Path}/bin/pulumi-tfgen-{ProviderName} {Lang} --out {OutputPath}/{Lang}/"
)

// BaseGenerateSdkCommand creates a shell command sequence for generating an SDK
// using the baseGenerateCmd template with the provided parameters
func BaseGenerateSdkCommand(providerName, path, outputPath, langString string) shell.ShellCommandSequence {
	cmd := baseGenerateSdkCmd
	cmd = strings.ReplaceAll(cmd, "{ProviderName}", providerName)
	cmd = strings.ReplaceAll(cmd, "{Path}", path)
	cmd = strings.ReplaceAll(cmd, "{OutputPath}", outputPath)
	cmd = strings.ReplaceAll(cmd, "{Lang}", langString)

	var result = shell.NewShellCommandSequence()
	result.Append(cmd)
	return result
}
