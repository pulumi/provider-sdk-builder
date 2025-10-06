package lang

import (
	"strings"
)

const (
	joinCmdLineEnding             = " && \\\n"
	pulumiPackageGenSdkCmd string = "pulumi package gen-sdk {Path} --language {Lang} --out {OutputPath} --version {Version}"
	//copyLicenseFile        string = "cp --out {OutputPath}"
	// TODO add license here
	// TODO generate readme.md here
)

// BaseGenerateSdkCommand creates a shell command for generating an SDK
// using the baseGenerateCmd template with the provided parameters
func BaseGenerateSdkCommand(schemaPath, outputPath, language, version string) []string {

	cmd := pulumiPackageGenSdkCmd

	cmd = strings.ReplaceAll(cmd, "{Path}", schemaPath)
	cmd = strings.ReplaceAll(cmd, "{Lang}", language)
	cmd = strings.ReplaceAll(cmd, "{OutputPath}", outputPath)
	cmd = strings.ReplaceAll(cmd, "{Version}", version)

	return []string{cmd}
}
