package lang

import (
	"strings"
)

type DotNet struct{}

func (p DotNet) String() string {
	return "dotnet"
}

func (p DotNet) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p DotNet) CompileSdkRecipe(outputPath string) []string {
	// Named individual commands for ease of comprehension
	const (
		makeNugetDir     = "mkdir -p nuget"
		cdToDotNetDirCmd = "cd {OutputPath}/dotnet/"
		buildDotNetCmd   = "dotnet build"
	)

	var compileDotNetRecipie = []string{
		makeNugetDir,
		cdToDotNetDirCmd,
		buildDotNetCmd,
	}

	compileDotNetCmd := strings.Join(compileDotNetRecipie, joinCmdLineEnding)
	compileDotNetCmd = strings.ReplaceAll(compileDotNetCmd, "{OutputPath}", outputPath)
	return []string{compileDotNetCmd}
}

func (p DotNet) PackageSdkRecipie() []string {
	return []string{}
}
