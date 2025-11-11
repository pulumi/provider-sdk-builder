package lang

import (
	"strings"
	"testing"
)

const (
	// expectedGoGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedGoGenerateCommandCount = 4

	// expectedGoCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. pack-sdk command
	expectedGoCompileCommandCount = 1

	// expectedGoInstallCommandCount is the number of commands returned by InstallSdkRecipe:
	// 0 (Go doesn't need install steps)
	expectedGoInstallCommandCount = 0
)

func TestGoLangGenerateSdkRecipe(t *testing.T) {
	goLang := GoLang{}
	schemaPath := "/path/to/schema.json"
	outputPath := "/output/dir"
	version := "1.0.0"
	providerPath := "/provider/path"

	result := goLang.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != expectedGoGenerateCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedGoGenerateCommandCount, len(result))
	}

	// Verify SDK generation command
	expectedSdkCmd := "pulumi package gen-sdk /path/to/schema.json --language go --out /output/dir --version 1.0.0"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/provider/path/README.md\" \"/output/dir/go/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/provider/path/LICENSE\" \"/output/dir/go/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
	}

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"1.0.0\" > \"/output/dir/go/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestGoLangCompileSdkRecipe(t *testing.T) {
	goLang := GoLang{}
	outputPath := "/test/output"
	providerPath := "/provider/path"

	result := goLang.CompileSdkRecipe(outputPath, providerPath)

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

func TestGoLangInstallSdkRecipe(t *testing.T) {
	goLang := GoLang{}
	outputPath := "/test/output"

	result := goLang.InstallSdkRecipe(outputPath)

	// Go should return empty slice (no install steps needed)
	if len(result) != expectedGoInstallCommandCount {
		t.Fatalf("expected %d commands for Go install (no-op), got %d", expectedGoInstallCommandCount, len(result))
	}
}
