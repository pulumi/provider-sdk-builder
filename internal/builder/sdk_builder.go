package builder

import (
	"github.com/pulumi/provider-sdk-builder/internal/lang"
)

type BuildParameters struct {
	ProviderName         string
	SchemaPath           string
	OutputPath           string
	RawRequestedLanguage string
}

type BuildInstructions struct {
	GenerateSdks      bool
	CompileSdks       bool
	PackageForRelease bool
}

func GenerateBuildCmds(params BuildParameters, instructions BuildInstructions) ([]string, error) {

	var result []string
	languages, err := lang.ParseRequestedLanguages(params.RawRequestedLanguage)
	if err != nil {
		return result, err
	}

	for _, chosenLanguage := range languages {
		// TODO to enable running this in parallel, we need to collect the commands for each language into a seprate list here
		if instructions.GenerateSdks {
			result = append(result, chosenLanguage.GenerateSdkRecipe(params.ProviderName, params.SchemaPath, params.OutputPath)...)
		}

		if instructions.CompileSdks {
			result = append(result, chosenLanguage.CompileSdkRecipe(params.OutputPath)...)
		}

		if instructions.PackageForRelease {
			result = append(result, chosenLanguage.PackageSdkRecipie()...)
		}
	}

	// TODO dispatch each language in its own thread

	return result, nil
}
