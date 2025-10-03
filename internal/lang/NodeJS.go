package lang

import (
	"strings"
)

type NodeJS struct{}

func (p NodeJS) String() string {
	return "nodejs"
}

func (p NodeJS) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p NodeJS) CompileSdkRecipe(outputPath string) []string {
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

func (p NodeJS) PackageSdkRecipie() []string {
	return []string{}
}
