package builder

import (
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/pulumi/provider-sdk-builder/internal/lang"
)

type BuildInstructions struct {
	GenerateSdks bool
	CompileSdks  bool
}

func BuildSDKs(params BuildParameters, instructions BuildInstructions, quiet bool, writer io.Writer) error {

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	var errorList []error

	// Spawn a goroutine for each language
	for _, chosenLanguage := range params.RequestedLanguages {
		waitGroup.Add(1)

		// Capture loop variable
		language := chosenLanguage

		go func() {
			defer waitGroup.Done()

			// Build SDK for this language
			err := BuildGivenLanguage(language, params, instructions, quiet, writer)
			if err != nil {
				mutex.Lock()
				errorList = append(errorList, err)
				mutex.Unlock()
			}
		}()
	}

	// Wait for all goroutines to complete
	waitGroup.Wait()

	// Check if any errors occurred
	if len(errorList) > 0 {
		// Combine all errors into one
		errorMessage := "SDK build errors occurred:"
		for _, err := range errorList {
			errorMessage += fmt.Sprintf("\n  - %v", err)
		}
		return errors.New(errorMessage)
	}

	return nil
}

func BuildGivenLanguage(language lang.Language, params BuildParameters, instructions BuildInstructions, quiet bool, writer io.Writer) error {
	var commands []string

	// Generate SDK commands if requested
	if instructions.GenerateSdks {
		languageSchemaPath, err := GetLanguageSchema(params.SchemaPath, language.String(), params.OutputPath)
		if err != nil {
			return fmt.Errorf("failed to get schema for %s: %w", language.String(), err)
		}
		commands = append(commands, language.GenerateSdkRecipe(languageSchemaPath, params.OutputPath, params.VersionString, params.ProviderPath)...)
	}

	// Compile SDK commands if requested
	if instructions.CompileSdks {
		commands = append(commands, language.CompileSdkRecipe(params.OutputPath, params.ProviderPath)...)
	}

	// Execute the commands for this language
	if len(commands) > 0 {
		if !quiet {
			fmt.Fprintf(writer, "\n=== Building SDK for %s ===\n", language.String())
		}

		err := ExecuteCommandSequence(commands, quiet, writer)
		if err != nil {
			return fmt.Errorf("failed to build %s SDK: %w", language.String(), err)
		}
	}

	return nil
}
