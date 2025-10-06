package lang

import (
	"strings"
	"testing"
)

func TestJavaGenerateSdkRecipe(t *testing.T) {
	java := Java{}
	schemaPath := "/schemas/gcp.json"
	outputPath := "/output/java"
	version := "3.2.1"
	providerPath := "/provider/path"

	result := java.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != 3 {
		t.Fatalf("expected 3 commands (SDK generation + README copy + LICENSE copy), got %d", len(result))
	}

	// Verify SDK generation command
	expectedSdkCmd := "pulumi package gen-sdk /schemas/gcp.json --language java --out /output/java --version 3.2.1"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/provider/path/README.md\" \"/output/java/java/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/provider/path/LICENSE\" \"/output/java/java/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
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