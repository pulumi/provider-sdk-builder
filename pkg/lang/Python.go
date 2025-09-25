package lang

import (
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

type Python struct{}

func (p Python) String() string {
	return "python"
}

func (p Python) GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Python) CompileSdkRecipe() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}

func (p Python) PackageSdkRecipie() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}
