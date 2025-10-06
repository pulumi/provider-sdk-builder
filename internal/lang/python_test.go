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

	result := python.GenerateSdkRecipe(schemaPath, outputPath, version)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /schemas/azure.json --language python --out /build/sdks --version 2.1.0"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
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