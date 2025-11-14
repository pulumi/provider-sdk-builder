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
		mkdirNuget              = "mkdir -p {InstallLocation}/nuget"
		checkNugetSourcesActive = "if dotnet nuget list source | grep \"{InstallLocation}/nuget\" | grep \"[Enabled]\"; then {CMD_TO_EXECUTE} ; fi"
		removePriorNugetSources = "dotnet nuget list source | grep Enabled | awk '{print $2}' | xargs -I % dotnet nuget remove source %"
		findAndCopyNupkg        = "find {SdkLocation} -name '*.nupkg' -print -exec cp -p \"{}\" {InstallLocation}/nuget \\;"
		checkAndAddSource       = "dotnet nuget add source \"{InstallLocation}/nuget\" --name \"{InstallLocation}/nuget\""
	)

	var installDotNetRecipe = []string{
		mkdirNuget,
		strings.ReplaceAll(checkNugetSourcesActive, "{CMD_TO_EXECUTE}", removePriorNugetSources),
		removePriorNugetSources,
		findAndCopyNupkg,
		checkAndAddSource,
	}

	installDotNetCmd := strings.Join(installDotNetRecipe, joinCmdLineEnding)
	installDotNetCmd = strings.ReplaceAll(installDotNetCmd, "{SdkLocation}", sdkLocation)
	installDotNetCmd = strings.ReplaceAll(installDotNetCmd, "{InstallLocation}", installLocation)
	return []string{installDotNetCmd}
}
