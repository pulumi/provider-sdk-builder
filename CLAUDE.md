# Provider SDK Builder - Developer Guide

## Project Overview

**Purpose**: CLI tool for generating Pulumi Provider SDKs across multiple programming languages.

**Supported Languages**: Go, Python, .NET (dotnet), NodeJS, Java

**Module**: `github.com/pulumi/provider-sdk-builder`

**Technology**: Go CLI using Cobra framework

## Architecture

The codebase follows a three-layer architecture:

```
CLI Layer (cmd/)           → User-facing commands
    ↓
Builder Layer (internal/builder/) → Command orchestration
    ↓
Language Layer (internal/lang/)   → Language-specific implementations
```

### Three-Phase SDK Build Process

1. **Generate** - Create SDK source code from schema
2. **Compile** - Package the SDK for distribution
3. **Install** - Install SDK locally for testing

## Commands

- `generate` - Generate SDK source code only
- `compile` - Compile/package SDK only
- `install` - Install SDK locally for testing only
- `build-sdks` - Run generate + compile phases

## Global Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--providerPath` | `-p` | `./` | Path to provider directory |
| `--language` | `-l` | `all` | Comma-separated language list or "all" |
| `--out` | `-o` | `{provider}/sdk` | SDK output directory |
| `--schema` | | `{provider}/provider/cmd/pulumi-resource-random/schema.json` | Path to schema file |
| `--version` | | `4.0.0-alpha.0+dev` | SDK version string |
| `--quiet` | `-q` | `false` | Hide command execution details |

## Code Organization

```
.
├── main.go                    # Entry point
├── cmd/                       # CLI commands (Cobra)
│   ├── root.go               # Root command + global flags
│   ├── generate_sdk.go       # Generate command
│   ├── compile_sdk.go        # Compile command
│   ├── install_sdk.go        # Install command
│   └── build_sdks.go         # Build-sdks command
└── internal/
    ├── builder/               # Core orchestration
    │   ├── parser.go         # Parse inputs → BuildParameters
    │   ├── sdk_builder.go    # Generate shell commands
    │   └── command_executor.go # Execute shell commands
    └── lang/                  # Language implementations
        ├── supported_language.go  # Language interface + registry
        ├── shared_steps.go        # Common generation logic
        ├── dotnet.go
        ├── golang.go
        ├── java.go
        ├── nodejs.go
        └── python.go
```

## Key Data Structures

### BuildParameters
```go
type BuildParameters struct {
    ProviderPath       string
    RequestedLanguages []lang.Language
    SchemaPath         string
    OutputPath         string
    VersionString      string
}
```

### BuildInstructions
```go
type BuildInstructions struct {
    GenerateSdks bool  // Run generate phase
    CompileSdks  bool  // Run compile phase
    InstallSdks  bool  // Run install phase
}
```

### Language Interface
```go
type Language interface {
    String() string
    GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string
    CompileSdkRecipe(outputPath, providerPath string) []string
    InstallSdkRecipe(outputPath string) []string
}
```

Each method returns a list of shell commands to execute for that phase.

## Flow

1. **User invokes command** (e.g., `provider-sdk-builder generate -l python`)
2. **Command handler** (cmd/*.go) calls `builder.ParseInputs()`
3. **Parser** creates BuildParameters with resolved paths and parsed languages
4. **Command handler** calls `builder.GenerateBuildCmds()` with BuildInstructions
5. **SDK Builder** iterates over languages, calling appropriate recipe methods
6. **Language implementations** return shell command strings
7. **Command executor** runs commands sequentially, capturing output
8. **Output** is printed to stdout

## Language Registration

Languages are registered in `internal/lang/supported_language.go`:

```go
var _languages = map[string]Language{
    "dotnet": dotNet,
    "go":     goLang,
    "java":   java,
    "nodejs": nodeJS,
    "python": python,
}
```

Use `"all"` to build for all registered languages.

## Testing

- Test files: `*_test.go` alongside implementation files
- Each language has its own test file (e.g., `dotnet_test.go`)
- Tests verify command generation, not execution

## Common Development Tasks

### Adding a New Language

1. Create `internal/lang/newlang.go`
2. Implement the `Language` interface
3. Register in `supported_language.go` map
4. Create `internal/lang/newlang_test.go`
5. Test all three recipe methods

### Modifying Command Generation

- **For all languages**: Edit `BaseGenerateSdkCommand()` in `shared_steps.go`
- **For specific language**: Edit that language's recipe methods

### Understanding Command Execution

- Commands execute sequentially via `ExecuteCommandSequence()`
- Each command runs in `/bin/sh -c`
- Execution stops on first error
- Use `--quiet` flag to suppress command echo

## Special Cases

### DotNet
- Compile phase creates a `nuget` directory at provider path
- Install phase has complex NuGet source management

### NodeJS
- Install phase creates yarn link for local testing

### Go, Java, Python
- No install steps (return empty slice)

## File Paths

- **Schema**: Defaults to `{providerPath}/provider/cmd/pulumi-resource-random/schema.json`
- **Output**: Defaults to `{providerPath}/sdk`
- **SDK location**: `{outputPath}/{language}/`
- **Version file**: `{outputPath}/{language}/version.txt`
