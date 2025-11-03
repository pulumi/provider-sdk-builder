package lang

import (
	"testing"
)

const (
	// expectedPythonGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedPythonGenerateCommandCount = 4

	// expectedPythonCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. pack-sdk command
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

	// Verify the command uses pulumi package pack-sdk
	cmd := result[0]
	expected := "cd /build/output/python && pulumi package pack-sdk python ."

	if cmd != expected {
		t.Errorf("expected command: %q\ngot:      %q", expected, cmd)
	}
}

func TestPythonInstallSdkRecipe(t *testing.T) {
	python := Python{}
	outputPath := "/build/output"

	result := python.InstallSdkRecipe(outputPath)

	// Python should return empty slice (no install steps needed)
	if len(result) != expectedPythonInstallCommandCount {
		t.Fatalf("expected %d commands for Python install (no-op), got %d", expectedPythonInstallCommandCount, len(result))
	}
}