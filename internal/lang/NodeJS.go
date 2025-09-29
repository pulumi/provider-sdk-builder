package lang

import ()

type NodeJS struct{}

func (p NodeJS) String() string {
	return "nodejs"
}

func (p NodeJS) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p NodeJS) CompileSdkRecipe() []string {
	return []string{}
}

func (p NodeJS) PackageSdkRecipie() []string {
	return []string{}
}
