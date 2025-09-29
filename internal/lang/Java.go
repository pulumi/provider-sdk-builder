package lang

import ()

type Java struct{}

func (p Java) String() string {
	return "java"
}

func (p Java) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Java) CompileSdkRecipe() []string {
	return []string{}
}

func (p Java) PackageSdkRecipie() []string {
	return []string{}
}
