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

func (l DotNet) CompileSdkRecipe(outputPath string) []string {
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

func (l DotNet) PackageSdkRecipie() []string {
	return []string{}
}
