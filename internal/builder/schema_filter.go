package builder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen/schemafilter"
)

func GetLanguageSchema(schemaPath, language, outputPath string) (string, error) {
	// Read the provider schema from file
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return "", err
	}

	languageSchemaBytes := schemafilter.FilterSchemaByLanguage(schemaBytes, language)

	languageSchemasDir := filepath.Join(outputPath, "language-schemas")
	err = os.MkdirAll(languageSchemasDir, 0o755)
	if err != nil {
		return "", err)
	}
	schemaFileName := fmt.Sprintf("%s-schema.json", language)
	languageSchemaPath := filepath.Join(languageSchemasDir, schemaFileName)

	err = os.WriteFile(languageSchemaPath, languageSchemaBytes, 0o600)
	if err != nil {
		return "", err
	}
	return languageSchemaPath, nil
}
