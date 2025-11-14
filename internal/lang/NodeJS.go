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
	// Named individual commands for ease of comprehension
	const (
		cdToNodeDirCmd             = "cd {OutputPath}/nodejs"
		installYarnCmd             = "yarn install"
		yarnRunTscCmd              = "yarn run tsc"
		copyPackageAndLockToBinCmd = "cp package.json yarn.lock ./bin/"
	)

	var compileNodeRecipie = []string{
		cdToNodeDirCmd,
		installYarnCmd,
		yarnRunTscCmd,
		copyPackageAndLockToBinCmd,
	}

	compileNodeCmd := strings.Join(compileNodeRecipie, joinCmdLineEnding)
	compileNodeCmd = strings.ReplaceAll(compileNodeCmd, "{OutputPath}", outputPath)
	return []string{compileNodeCmd}
}

func (l NodeJS) InstallSdkRecipe(outputPath string) []string {
	const installNodeCmd = "yarn link --cwd {OutputPath}/nodejs/bin"

	cmd := strings.ReplaceAll(installNodeCmd, "{OutputPath}", outputPath)
	return []string{cmd}
}
