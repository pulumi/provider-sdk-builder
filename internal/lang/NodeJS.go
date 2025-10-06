package lang

import (
	"strings"
)

type NodeJS struct{}

func (l NodeJS) String() string {
	return "nodejs"
}

func (l NodeJS) GenerateSdkRecipe(schemaPath, outputPath, version string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version)
}

func (l NodeJS) CompileSdkRecipe(outputPath string) []string {
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

func (l NodeJS) PackageSdkRecipie() []string {
	return []string{}
}
