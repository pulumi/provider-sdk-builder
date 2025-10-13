package lang

type Python struct{}

func (l Python) String() string {
	return "python"
}

func (l Python) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l Python) CompileSdkRecipe(outputPath, providerPath string) []string {
	return []string{"cd " + outputPath + "/python && pulumi package pack-sdk python ."}
}

func (l Python) InstallSdkRecipe(outputPath string) []string {
	// No install steps needed for Python
	return []string{}
}
