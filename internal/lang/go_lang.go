package lang

import "strings"

type GoLang struct{}

func (l GoLang) String() string {
	return "go"
}

func (l GoLang) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l GoLang) CompileSdkRecipe(outputPath, providerPath string) []string {
	//TODO once we stop checking in the other SDKS we can replace this recipe with the following one liner
	// return []string{"cd " + outputPath + " && pulumi package pack-sdk go ."}

	// Named individual commands for ease of comprehension
	const (
		cdToSdkDir        = "cd {OutputPath}/"
		goListAndBuildCmd = `go list "$(grep -e "^module" go.mod | cut -d ' ' -f 2)/go/..." | xargs -I {} bash -c 'go build {} && go clean -i {}'`
	)

	var compileGoRecipie = []string{
		cdToSdkDir,
		goListAndBuildCmd,
	}

	compileGoCmd := strings.Join(compileGoRecipie, joinCmdLineEnding)
	compileGoCmd = strings.ReplaceAll(compileGoCmd, "{OutputPath}", outputPath)
	return []string{compileGoCmd}
}

func (l GoLang) InstallSdkRecipe(sdkLocation, installLocation string) []string {
	// No install steps needed for Go
	return []string{}
}
