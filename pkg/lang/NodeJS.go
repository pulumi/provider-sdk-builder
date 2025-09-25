package lang

import (
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

type NodeJS struct{}

func (p NodeJS) String() string {
	return "nodejs"
}

func (p NodeJS) GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p NodeJS) CompileSdkRecipe() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}

func (p NodeJS) PackageSdkRecipie() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}
