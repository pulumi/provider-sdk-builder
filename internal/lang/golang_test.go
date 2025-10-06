package lang

import (
	"strings"
	"testing"
)

func TestGoLangGenerateSdkRecipe(t *testing.T) {
	goLang := GoLang{}
	schemaPath := "/path/to/schema.json"
	outputPath := "/output/dir"
	version := "1.0.0"

	result := goLang.GenerateSdkRecipe(schemaPath, outputPath, version)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /path/to/schema.json --language go --out /output/dir --version 1.0.0"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
	}
}

func TestGoLangCompileSdkRecipe(t *testing.T) {
	goLang := GoLang{}
	outputPath := "/test/output"

	result := goLang.CompileSdkRecipe(outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	// Verify the command contains the expected components
	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /test/output/") {
		t.Errorf("expected command to contain 'cd /test/output/', got: %q", cmd)
	}

	// Should contain the go list and build command
	if !strings.Contains(cmd, "go list") {
		t.Errorf("expected command to contain 'go list', got: %q", cmd)
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