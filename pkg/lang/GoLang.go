package lang

import (
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

type GoLang struct{}

func (p GoLang) String() string {
	return "go"
}

func (p GoLang) GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p GoLang) CompileSdkRecipe() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}

func (p GoLang) PackageSdkRecipie() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}
