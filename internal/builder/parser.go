package builder

import (
	"strings"

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
	schemaPathPattern = "provider/cmd/pulumi-resource-{ProviderName}/schema.json"
	defaultOutputPath = "sdk"
)

func ParseInputs(providerPath, providerName, rawRequestedLanguages, schemaPath, outputPath, versionString string) (BuildParameters, error) {

	languages, err := lang.ParseRequestedLanguages(rawRequestedLanguages)

	if err != nil {
		return BuildParameters{}, nil
	}

	if schemaPath == "" {
		// Substitute {ProviderName} in the pattern with the actual provider name
		pattern := strings.ReplaceAll(schemaPathPattern, "{ProviderName}", providerName)
		schemaPath = providerPath + pattern
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
