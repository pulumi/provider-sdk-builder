package lang

import (
	"strings"
	"testing"
)

func TestBaseGenerateSdkCommand(t *testing.T) {
	tests := []struct {
		name         string
		providerName string
		path         string
		outputPath   string
		langString   string
		expected     []string
	}{
		{
			name:         "with outputDir set",
			providerName: "aws",
			path:         "/path/to/schema.json",
			outputPath:   "/output/directory",
			langString:   "go",
			expected:     []string{"pulumi package gen-sdk /path/to/schema.json --language go --out /output/directory"},
		},
		{
			name:         "without outputDir",
			providerName: "azure",
			path:         "/path/to/azure-schema.json",
			outputPath:   "",
			langString:   "python",
			expected:     []string{"pulumi package gen-sdk /path/to/azure-schema.json --language python"},
		},
		{
			name:         "with different language and paths",
			providerName: "gcp",
			path:         "/schemas/gcp.json",
			outputPath:   "/build/sdks",
			langString:   "nodejs",
			expected:     []string{"pulumi package gen-sdk /schemas/gcp.json --language nodejs --out /build/sdks"},
		},
		{
			name:         "empty outputPath should not add --out flag",
			providerName: "k8s",
			path:         "/k8s-schema.json",
			outputPath:   "",
			langString:   "java",
			expected:     []string{"pulumi package gen-sdk /k8s-schema.json --language java"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BaseGenerateSdkCommand(tt.providerName, tt.path, tt.outputPath, tt.langString)

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

func TestBaseGenerateSdkCommandWithOutputDir(t *testing.T) {
	providerName := "test-provider"
	path := "/test/schema.json"
	outputPath := "/test/output"
	langString := "dotnet"

	result := BaseGenerateSdkCommand(providerName, path, outputPath, langString)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /test/schema.json --language dotnet --out /test/output"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
	}

	// Verify that the output path is included
	if !strings.Contains(result[0], "--out /test/output") {
		t.Errorf("expected command to contain '--out /test/output', got: %q", result[0])
	}
}

func TestBaseGenerateSdkCommandWithoutOutputDir(t *testing.T) {
	providerName := "test-provider"
	path := "/test/schema.json"
	outputPath := "" // Empty output path
	langString := "go"

	result := BaseGenerateSdkCommand(providerName, path, outputPath, langString)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /test/schema.json --language go"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
	}

	// Verify that the output path is NOT included
	if strings.Contains(result[0], "--out") {
		t.Errorf("expected command to NOT contain '--out' flag, got: %q", result[0])
	}
}