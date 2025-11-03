package lang

import (
	"testing"
)

const (
	// expectedJavaGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedJavaGenerateCommandCount = 4

	// expectedJavaCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. pack-sdk command
	expectedJavaCompileCommandCount = 1

	// expectedJavaInstallCommandCount is the number of commands returned by InstallSdkRecipe:
	// 0 (Java doesn't need install steps)
	expectedJavaInstallCommandCount = 0
)

func TestJavaGenerateSdkRecipe(t *testing.T) {
	java := Java{}
	schemaPath := "/schemas/gcp.json"
	outputPath := "/output/java"
	version := "3.2.1"
	providerPath := "/provider/path"

	result := java.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != expectedJavaGenerateCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedJavaGenerateCommandCount, len(result))
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

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"3.2.1\" > \"/output/java/java/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestJavaCompileSdkRecipe(t *testing.T) {
	java := Java{}
	outputPath := "/test/build"
	providerPath := "/provider/path"

	result := java.CompileSdkRecipe(outputPath, providerPath)

	if len(result) != expectedJavaCompileCommandCount {
		t.Fatalf("expected %d command, got %d", expectedJavaCompileCommandCount, len(result))
	}

	// Verify the command uses pulumi package pack-sdk
	cmd := result[0]
	expected := "cd /test/build/java && pulumi package pack-sdk java ."

	if cmd != expected {
		t.Errorf("expected command: %q\ngot:      %q", expected, cmd)
	}
}

func TestJavaInstallSdkRecipe(t *testing.T) {
	java := Java{}
	outputPath := "/test/build"

	result := java.InstallSdkRecipe(outputPath)

	// Java should return empty slice (no install steps needed)
	if len(result) != expectedJavaInstallCommandCount {
		t.Fatalf("expected %d commands for Java install (no-op), got %d", expectedJavaInstallCommandCount, len(result))
	}
}