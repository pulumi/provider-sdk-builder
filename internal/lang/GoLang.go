package lang

import ()

type GoLang struct{}

func (p GoLang) String() string {
	return "go"
}

func (p GoLang) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p GoLang) CompileSdkRecipe() []string {
	return []string{}
}

func (p GoLang) PackageSdkRecipie() []string {
	return []string{}
}
