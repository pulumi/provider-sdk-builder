package builder

import (
	"github.com/pulumi/provider-sdk-builder/pkg/lang"
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

func BuildSdks() error {
	return nil
}

func GenerateSdks(providerName, schemaPath, outputPath, rawLanguageString string) (shell.ShellCommandSequence, error) {

	languages, err := lang.ParseRequestedLanguages(rawLanguageString)
	if err != nil {
		return shell.ShellCommandSequence{}, err
	}
	return GenerateSdksCommands(providerName, schemaPath, outputPath, languages), nil
}

func GenerateSdksCommands(providerName, schemaPath, outputPath string, languages []lang.Language) shell.ShellCommandSequence {

	var result shell.ShellCommandSequence = shell.NewShellCommandSequence()

	for _, chosenLanguage := range languages {
		result.AppendAll(chosenLanguage.GenerateSdkRecipe(providerName, schemaPath, outputPath))
	}

	return result
}
