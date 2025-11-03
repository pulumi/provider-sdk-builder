package builder

import (
	"fmt"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen/schemafilter"
	"os"
	"path/filepath"
)

func GetLanguageSchema(schemaPath, language, outputPath string) (string, error) {
	// Read the provider schema from file
	schemaBytes, err := os.ReadFile(fmt.Sprintf(schemaPath))
	if err != nil {
		return "", err
	}

	languageSchemaBytes := schemafilter.FilterSchemaByLanguage(schemaBytes, language)

	schemaFileName := fmt.Sprintf("%s-schema.json", language)
	languageSchemaPath := filepath.Join(outputPath, schemaFileName)

	err = os.WriteFile(languageSchemaPath, languageSchemaBytes, 0o600)
	if err != nil {
		return "", err
	}
	return languageSchemaPath, nil
}
