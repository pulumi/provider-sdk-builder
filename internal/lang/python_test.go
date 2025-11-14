package lang

import (
	"strings"
	"testing"
)

const (
	// expectedPythonGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedPythonGenerateCommandCount = 4

	// expectedPythonCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. Compound command (cd + clean + copy + move + delete + venv + install + cd + build)
	expectedPythonCompileCommandCount = 1

	// expectedPythonInstallCommandCount is the number of commands returned by InstallSdkRecipe:
	// 0 (Python doesn't need install steps)
	expectedPythonInstallCommandCount = 0
)

func TestPythonGenerateSdkRecipe(t *testing.T) {
	python := Python{}
	schemaPath := "/schemas/azure.json"
	outputPath := "/build/sdks"
	version := "2.1.0"
	providerPath := "/provider/path"

	result := python.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != expectedPythonGenerateCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedPythonGenerateCommandCount, len(result))
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

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"2.1.0\" > \"/build/sdks/python/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestPythonCompileSdkRecipe(t *testing.T) {
	python := Python{}
	outputPath := "/build/output"
	providerPath := "/provider/path"

	result := python.CompileSdkRecipe(outputPath, providerPath)

	if len(result) != expectedPythonCompileCommandCount {
		t.Fatalf("expected %d command, got %d", expectedPythonCompileCommandCount, len(result))
	}

	// Verify the command contains the expected components
	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /build/output/python") {
		t.Errorf("expected command to contain 'cd /build/output/python', got: %q", cmd)
	}

	// Should contain clean bin command
	if !strings.Contains(cmd, "rm -rf ./bin/ ../python.bin/") {
		t.Errorf("expected command to contain 'rm -rf ./bin/ ../python.bin/', got: %q", cmd)
	}

	// Should contain copy command
	if !strings.Contains(cmd, "cp -R . ../python.bin") {
		t.Errorf("expected command to contain 'cp -R . ../python.bin', got: %q", cmd)
	}

	// Should contain move command
	if !strings.Contains(cmd, "mv ../python.bin ./bin") {
		t.Errorf("expected command to contain 'mv ../python.bin ./bin', got: %q", cmd)
	}

	// Should contain venv creation
	if !strings.Contains(cmd, "python3 -m venv venv") {
		t.Errorf("expected command to contain 'python3 -m venv venv', got: %q", cmd)
	}

	// Should contain build command
	if !strings.Contains(cmd, "../venv/bin/python -m build .") {
		t.Errorf("expected command to contain '../venv/bin/python -m build .', got: %q", cmd)
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
	sdkLocation := "/build/output"
	installLocation := ""

	result := python.InstallSdkRecipe(sdkLocation, installLocation)

	// Python should return empty slice (no install steps needed)
	if len(result) != expectedPythonInstallCommandCount {
		t.Fatalf("expected %d commands for Python install (no-op), got %d", expectedPythonInstallCommandCount, len(result))
	}
}