package lang

import ()

type Python struct{}

func (p Python) String() string {
	return "python"
}

func (p Python) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Python) CompileSdkRecipe() []string {
	return []string{}
}

func (p Python) PackageSdkRecipie() []string {
	return []string{}
}
