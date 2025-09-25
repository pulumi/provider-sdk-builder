package lang

import (
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

type Java struct{}

func (p Java) String() string {
	return "java"
}

func (p Java) GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Java) CompileSdkRecipe() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}

func (p Java) PackageSdkRecipie() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}
