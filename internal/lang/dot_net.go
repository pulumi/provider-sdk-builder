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

func (l DotNet) InstallSdkRecipe(sdkLocation, installLocation string) []string {
	// Named individual commands for ease of comprehension
	const (
		mkdirNuget               = "mkdir -p {InstallLocation}/nuget"
		findAndCopyNupkg         = "find {SdkLocation} -name '*.nupkg' -print -exec cp -p \"{}\" {InstallLocation}/nuget \\;"
		checkForLocalNugetSource = "if ! dotnet nuget list source | grep \"${InstallLocation}/nuget\"; then {CMD_TO_EXECUTE} ; fi"
		addNugetSource           = "dotnet nuget add source \"{InstallLocation}/nuget\" --name \"{InstallLocation}/nuget\""
	)

	var installDotNetRecipe = []string{
		mkdirNuget,
		findAndCopyNupkg,
		strings.ReplaceAll(checkForLocalNugetSource, "{CMD_TO_EXECUTE}", addNugetSource),
	}

	installDotNetCmd := strings.Join(installDotNetRecipe, joinCmdLineEnding)
	installDotNetCmd = strings.ReplaceAll(installDotNetCmd, "{SdkLocation}", sdkLocation)
	installDotNetCmd = strings.ReplaceAll(installDotNetCmd, "{InstallLocation}", installLocation)
	return []string{installDotNetCmd}
}
