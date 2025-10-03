package lang

import (
	"strings"
	"testing"
)

func TestNodeJSGenerateSdkRecipe(t *testing.T) {
	nodejs := NodeJS{}
	providerName := "k8s"
	path := "/k8s/schema.json"
	outputPath := "/sdk/output"

	result := nodejs.GenerateSdkRecipe(providerName, path, outputPath)

	if len(result) != 1 {
		t.Fatalf("expected 1 command, got %d", len(result))
	}

	expectedCmd := "pulumi package gen-sdk /k8s/schema.json --language nodejs --out /sdk/output"
	if result[0] != expectedCmd {
		t.Errorf("expected: %q\ngot:      %q", expectedCmd, result[0])
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