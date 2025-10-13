package lang

import (
	"strings"
	"testing"
)

func TestPythonGenerateSdkRecipe(t *testing.T) {
	python := Python{}
	schemaPath := "/schemas/azure.json"
	outputPath := "/build/sdks"
	version := "2.1.0"
	providerPath := "/provider/path"

	result := python.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != 3 {
		t.Fatalf("expected 3 commands (SDK generation + README copy + LICENSE copy), got %d", len(result))
	}

	// Verify SDK generation command
	expectedSdkCmd := "pulumi package gen-sdk /schemas/azure.json --language python --out /build/sdks --version 2.1.0"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/provider/path/README.md\" \"/build/sdks/python/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/provider/path/LICENSE\" \"/build/sdks/python/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
	}
}

func TestPythonCompileSdkRecipe(t *testing.T) {
	python := Python{}
	outputPath := "/build/output"

	result := python.CompileSdkRecipe(outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /build/output/python") {
		t.Errorf("expected command to contain 'cd /build/output/python', got: %q", cmd)
	}

	// Should contain key Python build commands
	if !strings.Contains(cmd, "python3 -m venv venv") {
		t.Errorf("expected command to contain 'python3 -m venv venv', got: %q", cmd)
	}

	if !strings.Contains(cmd, "python -m build .") {
		t.Errorf("expected command to contain 'python -m build .', got: %q", cmd)
	}

	// Should contain command joining
	if !strings.Contains(cmd, " && \\\n") {
		t.Errorf("expected command to contain command joiner ' && \\\\\\n', got: %q", cmd)
	}

	// Should have output path substituted (not contain template)
	if strings.Contains(cmd, "{OutputPath}") {
		t.Errorf("expected {OutputPath} to be substituted, but found it in: %q", cmd)
	}
}

func TestPythonInstallSdkRecipe(t *testing.T) {
	python := Python{}
	outputPath := "/build/output"

	result := python.InstallSdkRecipe(outputPath)

	// Python should return empty slice (no install steps needed)
	if len(result) != 0 {
		t.Fatalf("expected 0 commands for Python install (no-op), got %d", len(result))
	}
}