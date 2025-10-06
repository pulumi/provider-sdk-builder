package builder

import (
	"github.com/pulumi/provider-sdk-builder/internal/lang"
)

type BuildParameters struct {
	ProviderPath       string
	RequestedLanguages []lang.Language
	SchemaPath         string
	OutputPath         string
	VersionString      string
}

const (
	defaultSchemaPath = "provider/cmd/pulumi-resource-random/schema.json"
	defaultOutputPath = "sdk"
)

func ParseInputs(providerPath, rawRequestedLanguages, schemaPath, outputPath, versionString string) (BuildParameters, error) {

	languages, err := lang.ParseRequestedLanguages(rawRequestedLanguages)

	if err != nil {
		return BuildParameters{}, nil
	}

	if schemaPath == "" {
		schemaPath = providerPath + defaultSchemaPath
	}

	if outputPath == "" {
		outputPath = providerPath + defaultOutputPath
	}

	return BuildParameters{
		ProviderPath:       providerPath,
		RequestedLanguages: languages,
		SchemaPath:         schemaPath,
		OutputPath:         outputPath,
		VersionString:      versionString}, nil
}
