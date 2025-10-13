package lang

import (
	"strings"
)

type NodeJS struct{}

func (l NodeJS) String() string {
	return "nodejs"
}

func (l NodeJS) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l NodeJS) CompileSdkRecipe(outputPath, providerPath string) []string {
	return []string{"cd " + outputPath + "/nodejs && pulumi package pack-sdk nodejs ."}
}

func (l NodeJS) InstallSdkRecipe(outputPath string) []string {
	const installNodeCmd = "yarn link --cwd {OutputPath}/nodejs/bin"

	cmd := strings.ReplaceAll(installNodeCmd, "{OutputPath}", outputPath)
	return []string{cmd}
}
