package lang

import (
	"strings"
	"testing"
)

func TestDotNetGenerateSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	providerName := "random"
	path := "/random/schema.json"
	outputPath := "/dotnet/build"

	result := dotnet.GenerateSdkRecipe(providerName, path, outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /random/schema.json --language dotnet --out /dotnet/build"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
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