package lang

import (
	"strings"
)

type GoLang struct{}

func (l GoLang) String() string {
	return "go"
}

func (l GoLang) GenerateSdkRecipe(schemaPath, outputPath, version string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version)
}

func (l GoLang) CompileSdkRecipe(outputPath string) []string {
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

func (l GoLang) PackageSdkRecipie() []string {
	return []string{}
}
