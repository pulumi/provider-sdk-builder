package lang

import (
	"strings"
	"testing"
)

func TestDotNetGenerateSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	schemaPath := "/random/schema.json"
	outputPath := "/dotnet/build"
	version := "0.9.8"
	providerPath := "/provider/path"

	result := dotnet.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != 3 {
		t.Fatalf("expected 3 commands (SDK generation + README copy + LICENSE copy), got %d", len(result))
	}

	// Verify SDK generation command
	expectedSdkCmd := "pulumi package gen-sdk /random/schema.json --language dotnet --out /dotnet/build --version 0.9.8"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/provider/path/README.md\" \"/dotnet/build/dotnet/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/provider/path/LICENSE\" \"/dotnet/build/dotnet/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
	}
}

func TestDotNetCompileSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	outputPath := "/build/dotnet"

	result := dotnet.CompileSdkRecipe(outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /build/dotnet/dotnet/") {
		t.Errorf("expected command to contain 'cd /build/dotnet/dotnet/', got: %q", cmd)
	}

	// Should contain dotnet build command
	if !strings.Contains(cmd, "dotnet build") {
		t.Errorf("expected command to contain 'dotnet build', got: %q", cmd)
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