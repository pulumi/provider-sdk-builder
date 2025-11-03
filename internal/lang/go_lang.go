package lang

type GoLang struct{}

func (l GoLang) String() string {
	return "go"
}

func (l GoLang) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l GoLang) CompileSdkRecipe(outputPath, providerPath string) []string {
	return []string{"cd " + outputPath + " && pulumi package pack-sdk go ."}
}

func (l GoLang) InstallSdkRecipe(outputPath string) []string {
	// No install steps needed for Go
	return []string{}
}
