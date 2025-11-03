package lang

import (
	"strings"
	"testing"
)

const (
	// expectedDotNetGenerateCommandCount is the number of commands returned by GenerateSdkRecipe:
	// 1. SDK generation, 2. README copy, 3. LICENSE copy, 4. version.txt creation
	expectedDotNetGenerateCommandCount = 4

	// expectedDotNetCompileCommandCount is the number of commands returned by CompileSdkRecipe:
	// 1. mkdir for nuget directory, 2. pack-sdk command
	expectedDotNetCompileCommandCount = 2

	// expectedDotNetInstallCommandCount is the number of commands returned by InstallSdkRecipe:
	// 1. Complex command chain (mkdir + find + copy + nuget operations)
	expectedDotNetInstallCommandCount = 1
)

func TestDotNetGenerateSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	schemaPath := "/random/schema.json"
	outputPath := "/dotnet/build"
	version := "0.9.8"
	providerPath := "/provider/path"

	result := dotnet.GenerateSdkRecipe(schemaPath, outputPath, version, providerPath)

	if len(result) != expectedDotNetGenerateCommandCount {
		t.Fatalf("expected %d commands (SDK generation + README copy + LICENSE copy + version.txt), got %d", expectedDotNetGenerateCommandCount, len(result))
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

	// Verify version.txt creation command
	expectedVersionCmd := "echo \"0.9.8\" > \"/dotnet/build/dotnet/version.txt\""
	if result[3] != expectedVersionCmd {
		t.Errorf("expected version command: %q\ngot:      %q", expectedVersionCmd, result[3])
	}
}

func TestDotNetCompileSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	outputPath := "/build/dotnet"
	providerPath := "/provider/path"

	result := dotnet.CompileSdkRecipe(outputPath, providerPath)

	if len(result) != expectedDotNetCompileCommandCount {
		t.Fatalf("expected %d commands (mkdir nuget + pack-sdk), got %d", expectedDotNetCompileCommandCount, len(result))
	}

	// Verify mkdir command
	expectedMkdir := "mkdir -p /provider/path/nuget"
	if result[0] != expectedMkdir {
		t.Errorf("expected mkdir command: %q\ngot:      %q", expectedMkdir, result[0])
	}

	// Verify the pack-sdk command
	expectedPackSdk := "cd /build/dotnet/dotnet && pulumi package pack-sdk dotnet ."
	if result[1] != expectedPackSdk {
		t.Errorf("expected pack-sdk command: %q\ngot:      %q", expectedPackSdk, result[1])
	}
}

func TestDotNetInstallSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	outputPath := "/build/dotnet"

	result := dotnet.InstallSdkRecipe(outputPath)

	if len(result) != expectedDotNetInstallCommandCount {
		t.Fatalf("expected %d command, got %d", expectedDotNetInstallCommandCount, len(result))
	}

	cmd := result[0]

	// Should contain mkdir for nuget directory
	if !strings.Contains(cmd, "mkdir -p /build/dotnet/../nuget") {
		t.Errorf("expected command to contain 'mkdir -p /build/dotnet/../nuget', got: %q", cmd)
	}

	// Should contain find command to copy .nupkg files
	if !strings.Contains(cmd, "find /build/dotnet/dotnet/bin -name '*.nupkg'") {
		t.Errorf("expected command to contain find for .nupkg files, got: %q", cmd)
	}

	// Should contain copy command
	if !strings.Contains(cmd, "cp -p") {
		t.Errorf("expected command to contain 'cp -p', got: %q", cmd)
	}

	// Should contain dotnet nuget commands
	if !strings.Contains(cmd, "dotnet nuget list source") {
		t.Errorf("expected command to contain 'dotnet nuget list source', got: %q", cmd)
	}

	if !strings.Contains(cmd, "dotnet nuget add source") {
		t.Errorf("expected command to contain 'dotnet nuget add source', got: %q", cmd)
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