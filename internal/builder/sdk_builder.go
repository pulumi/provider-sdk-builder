package builder

import (
	"github.com/pulumi/provider-sdk-builder/internal/lang"
)

func BuildSdks() error {
	return nil
}

func GenerateSdksShellCommands(providerName, schemaPath, outputPath, rawLanguageString string) ([]string, error) {

	languages, err := lang.ParseRequestedLanguages(rawLanguageString)
	if err != nil {
		return []string{}, err
	}
	return buildGenerateShellCommands(providerName, schemaPath, outputPath, languages), nil
}

func buildGenerateShellCommands(providerName, schemaPath, outputPath string, languages []lang.Language) []string {

	var result []string

	for _, chosenLanguage := range languages {
		commands := chosenLanguage.GenerateSdkRecipe(providerName, schemaPath, outputPath)
		result = append(result, commands...)
	}

	return result
}
