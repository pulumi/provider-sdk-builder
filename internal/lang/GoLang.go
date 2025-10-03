package lang

import (
	"strings"
)

type GoLang struct{}

func (p GoLang) String() string {
	return "go"
}

func (p GoLang) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p GoLang) CompileSdkRecipe(outputPath string) []string {
	// Named individual commands for ease of comprehension
	const (
		cdToSdkDir        = "cd {OutputPath}/"
		goListAndBuildCmd = "go list \"$(grep -e \"^module\" go.mod | cut -d ' ' -f 2)/go/...\" | xargs -I {} bash -c 'go build {} && go clean -i {}'"
	)

	var compileGoRecipie = []string{
		cdToSdkDir,
		goListAndBuildCmd,
	}

	compileGoCmd := strings.Join(compileGoRecipie, joinCmdLineEnding)
	compileGoCmd = strings.ReplaceAll(compileGoCmd, "{OutputPath}", outputPath)
	return []string{compileGoCmd}
}

func (p GoLang) PackageSdkRecipie() []string {
	return []string{}
}
