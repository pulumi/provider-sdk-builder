package lang

import (
	"strings"
	"testing"
)

func TestNodeJSGenerateSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	schemaPath := "/k8s/schema.json"
	outputPath := "/sdk/output"
	version := "4.5.6"
	providerPath := "/provider/path"

	result := nodejs.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != 3 {
		t.Fatalf("expected 3 commands (SDK generation + README copy + LICENSE copy), got %d", len(result))
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
}

func TestNodeJSCompileSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	outputPath := "/project/output"

	result := nodejs.CompileSdkRecipe(outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	cmd := result[0]

	// Should contain the cd command with output path
	if !strings.Contains(cmd, "cd /project/output/nodejs") {
		t.Errorf("expected command to contain 'cd /project/output/nodejs', got: %q", cmd)
	}

	// Should contain yarn commands
	if !strings.Contains(cmd, "yarn install") {
		t.Errorf("expected command to contain 'yarn install', got: %q", cmd)
	}

	if !strings.Contains(cmd, "yarn run tsc") {
		t.Errorf("expected command to contain 'yarn run tsc', got: %q", cmd)
	}

	// Should contain file copy command
	if !strings.Contains(cmd, "cp package.json yarn.lock ./bin/") {
		t.Errorf("expected command to contain 'cp package.json yarn.lock ./bin/', got: %q", cmd)
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