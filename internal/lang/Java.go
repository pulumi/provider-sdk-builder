package lang

import (
	"strings"
)

type Java struct{}

func (p Java) String() string {
	return "java"
}

func (p Java) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Java) CompileSdkRecipe(outputPath string) []string {

	// Named individual commands for ease of comprehension
	const (
		cdToJavaDir      = "cd {OutputPath}/java/"
		gradleBuildCmd   = "gradle --console=plain build"
		gradleJavaDocCmd = "gradle --console=plain javadoc"
	)

	var compileJavaRecipie = []string{
		cdToJavaDir,
		gradleBuildCmd,
		gradleJavaDocCmd,
	}

	compileJavaCmd := strings.Join(compileJavaRecipie, joinCmdLineEnding)
	compileJavaCmd = strings.ReplaceAll(compileJavaCmd, "{OutputPath}", outputPath)
	return []string{compileJavaCmd}
}

func (p Java) PackageSdkRecipie() []string {
	return []string{}
}
