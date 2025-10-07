package builder

import (
	"errors"
)

type BuildInstructions struct {
	GenerateSdks bool
	CompileSdks  bool
	InstallSdks  bool
}

func GenerateBuildCmds(params BuildParameters, instructions BuildInstructions) ([]string, error) {

	var result []string

	for _, chosenLanguage := range params.RequestedLanguages {
		// TODO to enable running this in parallel, we need to collect the commands for each language into a seprate list here
		if instructions.GenerateSdks {
			result = append(result, chosenLanguage.GenerateSdkRecipe(params.SchemaPath, params.OutputPath, params.VersionString, params.ProviderPath)...)
		}

		if instructions.CompileSdks {
			result = append(result, chosenLanguage.CompileSdkRecipe(params.OutputPath)...)
		}

		if instructions.InstallSdks {
			result = append(result, chosenLanguage.InstallSdkRecipe(params.OutputPath)...)
		}
	}

	if len(result) == 0 {
		return result, errors.New("Empty list of commands generated from GenerateBuildCommand, aborting build")
	}

	return result, nil
}
