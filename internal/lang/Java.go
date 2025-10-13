package lang

type Java struct{}

func (l Java) String() string {
	return "java"
}

func (l Java) GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string {
	return BaseGenerateSdkCommand(schemaPath, outputPath, l.String(), version, providerPath)
}

func (l Java) CompileSdkRecipe(outputPath, providerPath string) []string {
	return []string{"cd " + outputPath + "/java && pulumi package pack-sdk java ."}
}

func (l Java) InstallSdkRecipe(outputPath string) []string {
	// No install steps needed for Java
	return []string{}
}
