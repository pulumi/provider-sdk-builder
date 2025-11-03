package lang

import (
	"strings"
)

type DotNet struct{}

func (l DotNet) String() string {
	return "dotnet"
}

func (l DotNet) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l DotNet) CompileSdkRecipe(outputPath, providerPath string) []string {
	mkdirCmd := "mkdir -p " + providerPath + "/nuget"
	compileCmd := "cd " + outputPath + "/dotnet && pulumi package pack-sdk dotnet ."
	return []string{mkdirCmd, compileCmd}
}

func (l DotNet) InstallSdkRecipe(outputPath string) []string {
	// Named individual commands for ease of comprehension
	const (
		mkdirNuget        = "mkdir -p {OutputPath}/../nuget"
		findAndCopyNupkg  = "find {OutputPath}/dotnet/bin -name '*.nupkg' -print -exec cp -p \"{{}}\" {OutputPath}/../nuget \\;"
		checkAndAddSource = "dotnet nuget list source | grep \"{OutputPath}/../nuget\" || dotnet nuget add source \"{OutputPath}/../nuget\" --name \"{OutputPath}/../nuget\""
	)

	var installDotNetRecipe = []string{
		mkdirNuget,
		findAndCopyNupkg,
		checkAndAddSource,
	}

	installDotNetCmd := strings.Join(installDotNetRecipe, joinCmdLineEnding)
	installDotNetCmd = strings.ReplaceAll(installDotNetCmd, "{OutputPath}", outputPath)
	return []string{installDotNetCmd}
}
