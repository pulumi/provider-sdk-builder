package lang

import ()

type DotNet struct{}

func (p DotNet) String() string {
	return "dotnet"
}

func (p DotNet) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p DotNet) CompileSdkRecipe() []string {
	return []string{}
}

func (p DotNet) PackageSdkRecipie() []string {
	return []string{}
}
