package lang

import (
	"path/filepath"
	"strings"
)

const (
	joinCmdLineEnding             = " && \\\n"
	pulumiPackageGenSdkCmd string = "pulumi package gen-sdk {Path} --language {Lang} --out {OutputPath} --version {Version}"
)

// BaseGenerateSdkCommand creates shell commands for generating an SDK and copying README and LICENSE files
func BaseGenerateSdkCommand(schemaPath, outputPath, language, version, providerPath string) []string {

	// Generate SDK command
	sdkCmd := pulumiPackageGenSdkCmd
	sdkCmd = strings.ReplaceAll(sdkCmd, "{Path}", schemaPath)
	sdkCmd = strings.ReplaceAll(sdkCmd, "{Lang}", language)
	sdkCmd = strings.ReplaceAll(sdkCmd, "{OutputPath}", outputPath)
	sdkCmd = strings.ReplaceAll(sdkCmd, "{Version}", version)

	// Generate file copy commands
	readmeCmd := generateCopyCommand(providerPath, "README.md", outputPath, language, "README.md")
	licenseCmd := generateCopyCommand(providerPath, "LICENSE", outputPath, language, "LICENSE")

	// Generate version.txt file command
	versionFilePath := filepath.Join(outputPath, language, "version.txt")
	versionCmd := "echo \"" + version + "\" > \"" + versionFilePath + "\""

	return []string{sdkCmd, readmeCmd, licenseCmd, versionCmd}
}

// generateCopyCommand creates a shell command to copy a file from source to destination
func generateCopyCommand(sourcePath, sourceFile, outputPath, language, destFile string) string {
	source := filepath.Join(sourcePath, sourceFile)
	dest := filepath.Join(outputPath, language, destFile)
	return "cp -f \"" + source + "\" \"" + dest + "\""
}
