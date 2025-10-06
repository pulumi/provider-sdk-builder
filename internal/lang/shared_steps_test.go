package lang

import (
	"strings"
	"testing"
)

func TestBaseGenerateSdkCommand(t *testing.T) {
	tests := []struct {
		name       string
		schemaPath string
		outputPath string
		language   string
		version    string
		expected   []string
	}{
		{
			name:       "go language with output path",
			schemaPath: "/path/to/schema.json",
			outputPath: "/output/directory",
			language:   "go",
			version:    "1.2.3",
			expected:   []string{"pulumi package gen-sdk /path/to/schema.json --language go --out /output/directory --version 1.2.3"},
		},
		{
			name:       "python language",
			schemaPath: "/schemas/multi.json",
			outputPath: "/build/sdks",
			language:   "python",
			version:    "2.0.1",
			expected:   []string{"pulumi package gen-sdk /schemas/multi.json --language python --out /build/sdks --version 2.0.1"},
		},
		{
			name:       "java language with different path",
			schemaPath: "/azure-schema.json",
			outputPath: "/azure/sdk",
			language:   "java",
			version:    "3.1.0",
			expected:   []string{"pulumi package gen-sdk /azure-schema.json --language java --out /azure/sdk --version 3.1.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BaseGenerateSdkCommand(tt.schemaPath, tt.outputPath, tt.language, tt.version)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d commands, got %d", len(tt.expected), len(result))
				return
			}

			for i, expectedCmd := range tt.expected {
				if result[i] != expectedCmd {
					t.Errorf("command %d mismatch:\nexpected: %q\ngot:      %q", i, expectedCmd, result[i])
				}
			}
		})
	}
}

func TestBaseGenerateSdkCommandDotNet(t *testing.T) {
	schemaPath := "/test/schema.json"
	outputPath := "/test/output"
	language := "dotnet"
	version := "1.5.2"

	result := BaseGenerateSdkCommand(schemaPath, outputPath, language, version)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /test/schema.json --language dotnet --out /test/output --version 1.5.2"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
	}

	// Verify that the language string is correctly substituted
	if !strings.Contains(result[0], "--language dotnet") {
		t.Errorf("expected command to contain '--language dotnet', got: %q", result[0])
	}
}