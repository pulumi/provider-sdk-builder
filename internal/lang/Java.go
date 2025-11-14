package lang

import (
	"strings"
)

type Java struct{}

func (l Java) String() string {
	return "java"
}

func (l Java) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l Java) CompileSdkRecipe(outputPath, providerPath string) []string {

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

func (l Java) InstallSdkRecipe(outputPath string) []string {
	// No install steps needed for Java
	return []string{}
}
