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
	// 1. Compound command (mkdir + cd + dotnet build)
	expectedDotNetCompileCommandCount = 1

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
		t.Fatalf("expected %d command, got %d", expectedDotNetCompileCommandCount, len(result))
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

func TestDotNetInstallSdkRecipe(t *testing.T) {
	dotnet := DotNet{}
	sdkLocation := "/build/sdk"
	installLocation := "/usr/local"

	result := dotnet.InstallSdkRecipe(sdkLocation, installLocation)

	if len(result) != expectedDotNetInstallCommandCount {
		t.Fatalf("expected %d command, got %d", expectedDotNetInstallCommandCount, len(result))
	}

	cmd := result[0]

	// Should contain mkdir for nuget directory at install location
	if !strings.Contains(cmd, "mkdir -p /usr/local/nuget") {
		t.Errorf("expected command to contain 'mkdir -p /usr/local/nuget', got: %q", cmd)
	}

	// Should contain find command to search sdk location for .nupkg files
	if !strings.Contains(cmd, "find /build/sdk -name '*.nupkg'") {
		t.Errorf("expected command to contain 'find /build/sdk -name '*.nupkg'', got: %q", cmd)
	}

	// Should copy to install location
	if !strings.Contains(cmd, "cp -p") && !strings.Contains(cmd, "/usr/local/nuget") {
		t.Errorf("expected command to copy to install location, got: %q", cmd)
	}

	// Should contain dotnet nuget add source command
	if !strings.Contains(cmd, "dotnet nuget add source") {
		t.Errorf("expected command to contain 'dotnet nuget add source', got: %q", cmd)
	}

	// Should reference install location in nuget source
	if !strings.Contains(cmd, "\"/usr/local/nuget\"") {
		t.Errorf("expected command to reference install location in nuget source, got: %q", cmd)
	}

	// Should contain command joining
	if !strings.Contains(cmd, " && \\\n") {
		t.Errorf("expected command to contain command joiner ' && \\\\\\n', got: %q", cmd)
	}

	// Should have templates substituted (not contain template variables)
	if strings.Contains(cmd, "{SdkLocation}") {
		t.Errorf("expected {SdkLocation} to be substituted, but found it in: %q", cmd)
	}

	if strings.Contains(cmd, "{InstallLocation}") {
		t.Errorf("expected {InstallLocation} to be substituted, but found it in: %q", cmd)
	}
}