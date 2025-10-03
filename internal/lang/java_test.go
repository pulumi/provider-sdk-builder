package lang

import (
	"strings"
	"testing"
)

func TestJavaGenerateSdkRecipe(t *testing.T) {
	java := Java{}
	providerName := "gcp"
	path := "/schemas/gcp.json"
	outputPath := "/output/java"

	result := java.GenerateSdkRecipe(providerName, path, outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /schemas/gcp.json --language java --out /output/java"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
	}
}

func TestJavaCompileSdkRecipe(t *testing.T) {
	java := Java{}
	outputPath := "/test/build"

	result := java.CompileSdkRecipe(outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /test/build/java/") {
		t.Errorf("expected command to contain 'cd /test/build/java/', got: %q", cmd)
	}

	// Should contain gradle build commands
	if !strings.Contains(cmd, "gradle --console=plain build") {
		t.Errorf("expected command to contain 'gradle --console=plain build', got: %q", cmd)
	}

	if !strings.Contains(cmd, "gradle --console=plain javadoc") {
		t.Errorf("expected command to contain 'gradle --console=plain javadoc', got: %q", cmd)
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