package builder

import (
	"testing"
)

func TestParseInputsProviderNameSubstitution(t *testing.T) {
	tests := []struct {
		name                string
		providerPath        string
		providerName        string
		rawRequestedLangs   string
		schemaPath          string
		outputPath          string
		versionString       string
		expectedSchemaPath  string
	}{
		{
			name:               "provider name substitution",
			providerPath:       "/test/path/",
			providerName:       "random",
			rawRequestedLangs:  "go",
			schemaPath:         "",
			outputPath:         "",
			versionString:      "1.0.0",
			expectedSchemaPath: "/test/path/provider/cmd/pulumi-resource-random/schema.json",
		},
		{
			name:               "provider name substitution with different name",
			providerPath:       "/test/path/",
			providerName:       "aws",
			rawRequestedLangs:  "python",
			schemaPath:         "",
			outputPath:         "",
			versionString:      "1.0.0",
			expectedSchemaPath: "/test/path/provider/cmd/pulumi-resource-aws/schema.json",
		},
		{
			name:               "explicit schema path overrides substitution",
			providerPath:       "/test/path/",
			providerName:       "random",
			rawRequestedLangs:  "go",
			schemaPath:         "/custom/path/schema.json",
			outputPath:         "",
			versionString:      "1.0.0",
			expectedSchemaPath: "/custom/path/schema.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := ParseInputs(
				tt.providerPath,
				tt.providerName,
				tt.rawRequestedLangs,
				tt.schemaPath,
				tt.outputPath,
				tt.versionString,
			)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if params.SchemaPath != tt.expectedSchemaPath {
				t.Errorf("expected SchemaPath to be %q, got %q", tt.expectedSchemaPath, params.SchemaPath)
			}

			// Verify other fields are set correctly
			if params.ProviderPath != tt.providerPath {
				t.Errorf("expected ProviderPath to be %q, got %q", tt.providerPath, params.ProviderPath)
			}

			if params.VersionString != tt.versionString {
				t.Errorf("expected VersionString to be %q, got %q", tt.versionString, params.VersionString)
			}
		})
	}
}

func TestParseInputsOutputPathDefault(t *testing.T) {
	params, err := ParseInputs(
		"/test/path/",
		"random",
		"go",
		"",
		"", // empty output path
		"1.0.0",
	)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedOutputPath := "/test/path/sdk"
	if params.OutputPath != expectedOutputPath {
		t.Errorf("expected OutputPath to be %q, got %q", expectedOutputPath, params.OutputPath)
	}
}

func TestParseInstallInputs(t *testing.T) {
	tests := []struct {
		name                  string
		rawRequestedLanguages string
		sdkLocation           string
		installLocation       string
		expectedLanguages     int
		expectError           bool
	}{
		{
			name:                  "single language",
			rawRequestedLanguages: "go",
			sdkLocation:           "/test/sdk",
			installLocation:       "/test/install",
			expectedLanguages:     1,
			expectError:           false,
		},
		{
			name:                  "multiple languages",
			rawRequestedLanguages: "go,python,nodejs",
			sdkLocation:           "/test/sdk",
			installLocation:       "/test/install",
			expectedLanguages:     3,
			expectError:           false,
		},
		{
			name:                  "all languages",
			rawRequestedLanguages: "all",
			sdkLocation:           "/test/sdk",
			installLocation:       "/test/install",
			expectedLanguages:     5,
			expectError:           false,
		},
		{
			name:                  "invalid language",
			rawRequestedLanguages: "invalid",
			sdkLocation:           "/test/sdk",
			installLocation:       "/test/install",
			expectedLanguages:     0,
			expectError:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := ParseInstallInputs(
				tt.rawRequestedLanguages,
				tt.sdkLocation,
				tt.installLocation,
			)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(params.RequestedLanguages) != tt.expectedLanguages {
				t.Errorf("expected %d languages, got %d", tt.expectedLanguages, len(params.RequestedLanguages))
			}

			if params.SdkLocation != tt.sdkLocation {
				t.Errorf("expected SdkLocation to be %q, got %q", tt.sdkLocation, params.SdkLocation)
			}

			if params.InstallLocation != tt.installLocation {
				t.Errorf("expected InstallLocation to be %q, got %q", tt.installLocation, params.InstallLocation)
			}
		})
	}
}
