package lang

import (
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

type DotNet struct{}

func (p DotNet) String() string {
	return "dotnet"
}

func (p DotNet) GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p DotNet) CompileSdkRecipe() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}

func (p DotNet) PackageSdkRecipie() shell.ShellCommandSequence {
	return shell.ShellCommandSequence{}
}
