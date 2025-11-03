package lang

import (
	"strings"
	"testing"
)

const (
	// expectedNodeJSGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedNodeJSGenerateCommandCount = 4

	// expectedNodeJSCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. pack-sdk command
	expectedNodeJSCompileCommandCount = 1

	// expectedNodeJSInstallCommandCount is the number of commands returned by InstallSdkRecipe:
	// 1. yarn link command
	expectedNodeJSInstallCommandCount = 1
)

func TestNodeJSGenerateSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	schemaPath := "/k8s/schema.json"
	outputPath := "/sdk/output"
	version := "4.5.6"
	providerPath := "/provider/path"

	result := nodejs.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != expectedNodeJSGenerateCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedNodeJSGenerateCommandCount, len(result))
	}

	// Verify SDK generation command
	expectedSdkCmd := "pulumi package gen-sdk /k8s/schema.json --language nodejs --out /sdk/output --version 4.5.6"
	if result[0] != expectedSdkCmd {
		t.Errorf("expected SDK command: %q\ngot:      %q", expectedSdkCmd, result[0])
	}

	// Verify README copy command
	expectedReadmeCmd := "cp -f \"/provider/path/README.md\" \"/sdk/output/nodejs/README.md\""
	if result[1] != expectedReadmeCmd {
		t.Errorf("expected README command: %q\ngot:      %q", expectedReadmeCmd, result[1])
	}

	// Verify LICENSE copy command
	expectedLicenseCmd := "cp -f \"/provider/path/LICENSE\" \"/sdk/output/nodejs/LICENSE\""
	if result[2] != expectedLicenseCmd {
		t.Errorf("expected LICENSE command: %q\ngot:      %q", expectedLicenseCmd, result[2])
	}

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"4.5.6\" > \"/sdk/output/nodejs/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestNodeJSCompileSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	outputPath := "/project/output"
	providerPath := "/provider/path"

	result := nodejs.CompileSdkRecipe(outputPath, providerPath)

	if len(result) != expectedNodeJSCompileCommandCount {
		t.Fatalf("expected %d command, got %d", expectedNodeJSCompileCommandCount, len(result))
	}

	// Verify the command uses pulumi package pack-sdk
	cmd := result[0]
	expected := "cd /project/output/nodejs && pulumi package pack-sdk nodejs ."

	if cmd != expected {
		t.Errorf("expected command: %q\ngot:      %q", expected, cmd)
	}
}

func TestNodeJSInstallSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	outputPath := "/project/output"

	result := nodejs.InstallSdkRecipe(outputPath)

	if len(result) != expectedNodeJSInstallCommandCount {
		t.Fatalf("expected %d command, got %d", expectedNodeJSInstallCommandCount, len(result))
	}

	cmd := result[0]

	// Should contain yarn link command with correct path
	expected := "yarn link --cwd /project/output/nodejs/bin"
	if cmd != expected {
		t.Errorf("expected command: %q\ngot:      %q", expected, cmd)
	}

	// Should have output path substituted (not contain template)
	if strings.Contains(cmd, "{OutputPath}") {
		t.Errorf("expected {OutputPath} to be substituted, but found it in: %q", cmd)
	}
}