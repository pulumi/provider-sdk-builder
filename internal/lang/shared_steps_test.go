package lang

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	// expectedGenerateSdkCommandCount is the number of commands returned by BaseGenerateSdkCommand:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedGenerateSdkCommandCount = 4
)

func TestBaseGenerateSdkCommand(t *testing.T) {
	tests := []struct {
		name         string
		schemaPath   string
		outputPath   string
		language     string
		version      string
		providerPath string
		expectedSdk  string
	}{
		{
			name:         "go language with output path",
			schemaPath:   "/path/to/schema.json",
			outputPath:   "/output/directory",
			language:     "go",
			version:      "1.2.3",
			providerPath: "/provider/path",
			expectedSdk:  "pulumi package gen-sdk /path/to/schema.json --language go --out /output/directory --version 1.2.3",
		},
		{
			name:         "python language",
			schemaPath:   "/schemas/multi.json",
			outputPath:   "/build/sdks",
			language:     "python",
			version:      "2.0.1",
			providerPath: "/provider/src",
			expectedSdk:  "pulumi package gen-sdk /schemas/multi.json --language python --out /build/sdks --version 2.0.1",
		},
		{
			name:         "java language with different path",
			schemaPath:   "/azure-schema.json",
			outputPath:   "/azure/sdk",
			language:     "java",
			version:      "3.1.0",
			providerPath: "/azure/provider",
			expectedSdk:  "pulumi package gen-sdk /azure-schema.json --language java --out /azure/sdk --version 3.1.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BaseGenerateSdkCommand(tt.schemaPath, tt.outputPath, tt.language, tt.version, tt.providerPath)

			if len(result) != expectedGenerateSdkCommandCount {
				t.Errorf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedGenerateSdkCommandCount, len(result))
				return
			}

			// Check SDK generation command
			if result[0] != tt.expectedSdk {
				t.Errorf("SDK command mismatch:\nexpected: %q\ngot:      %q", tt.expectedSdk, result[0])
			}

			// Check README copy command
			expectedReadme := "cp -f \"" + tt.providerPath + "/README.md\" \"" + tt.outputPath + "/" + tt.language + "/README.md\""
			if result[1] != expectedReadme {
				t.Errorf("README command mismatch:\nexpected: %q\ngot:      %q", expectedReadme, result[1])
			}

			// Check LICENSE copy command
			expectedLicense := "cp -f \"" + tt.providerPath + "/LICENSE\" \"" + tt.outputPath + "/" + tt.language + "/LICENSE\""
			if result[2] != expectedLicense {
				t.Errorf("LICENSE command mismatch:\nexpected: %q\ngot:      %q", expectedLicense, result[2])
			}

			// Check version.txt creation command
			expectedVersion := "echo \"" + tt.version + "\" > \"" + tt.outputPath + "/" + tt.language + "/version.txt\""
			if result[3] != expectedVersion {
				t.Errorf("version.txt command mismatch:\nexpected: %q\ngot:      %q", expectedVersion, result[3])
			}
		})
	}
}

func TestBaseGenerateSdkCommandDotNet(t *testing.T) {
	schemaPath := "/test/schema.json"
	outputPath := "/test/output"
	language := "dotnet"
	version := "1.5.2"
	providerPath := "/test/provider"

	result := BaseGenerateSdkCommand(schemaPath, outputPath, language, version, providerPath)

	if len(result) != expectedGenerateSdkCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedGenerateSdkCommandCount, len(result))
	}

	expectedSdkCmd := "pulumi package gen-sdk /test/schema.json --language dotnet --out /test/output --version 1.5.2"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify that the language string is correctly substituted
	if !strings.Contains(result[0], "--language dotnet") {
		t.Errorf("expected command to contain '--language dotnet', got: %q", result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/test/provider/README.md\" \"/test/output/dotnet/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/test/provider/LICENSE\" \"/test/output/dotnet/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
	}

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"1.5.2\" > \"/test/output/dotnet/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestBaseGenerateSdkCommandWithOverlays(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "overlay-test-*")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name            string
		setup           func() string // returns providerPath
		expectedHasFlag bool
		description     string
	}{
		{
			name: "overlays directory exists",
			setup: func() string {
				providerPath := filepath.Join(tmpDir, "provider1")
				os.MkdirAll(providerPath, 0755)
				overlayDir := filepath.Join(providerPath, "overlays")
				os.MkdirAll(overlayDir, 0755)
				return providerPath
			},
			expectedHasFlag: true,
			description:     "should add --overlays flag when overlays directory exists",
		},
		{
			name: "overlays directory does not exist",
			setup: func() string {
				providerPath := filepath.Join(tmpDir, "provider2")
				os.MkdirAll(providerPath, 0755)
				return providerPath
			},
			expectedHasFlag: false,
			description:     "should not add --overlays flag when overlays directory does not exist",
		},
		{
			name: "overlays is a file, not a directory",
			setup: func() string {
				providerPath := filepath.Join(tmpDir, "provider3")
				os.MkdirAll(providerPath, 0755)
				overlayFile := filepath.Join(providerPath, "overlays")
				os.WriteFile(overlayFile, []byte("not a directory"), 0644)
				return providerPath
			},
			expectedHasFlag: false,
			description:     "should not add --overlays flag when 'overlays' is a file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerPath := tt.setup()
			schemaPath := "/test/schema.json"
			outputPath := "/test/output"
			language := "nodejs"
			version := "1.0.0"

			result := BaseGenerateSdkCommand(schemaPath, outputPath, language, version, providerPath)

			if len(result) == 0 {
				t.Fatal("expected at least one command")
			}

			sdkCmd := result[0]
			fmt.Println(sdkCmd)
			hasOverlaysFlag := strings.Contains(sdkCmd, "--overlays")

			if hasOverlaysFlag != tt.expectedHasFlag {
				t.Errorf("%s\nExpected --overlays flag: %v, got: %v\nCommand: %q", tt.description, tt.expectedHasFlag, hasOverlaysFlag, sdkCmd)
			}

			if tt.expectedHasFlag {
				// Verify the overlay path is included
				expectedPath := filepath.Join(providerPath, "overlays")
				if !strings.Contains(sdkCmd, expectedPath) {
					t.Errorf("expected command to contain overlay path %q, got: %q", expectedPath, sdkCmd)
				}
			}
		})
	}
}
