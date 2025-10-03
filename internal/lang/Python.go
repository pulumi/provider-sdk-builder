package lang

import (
	"strings"
)

type Python struct{}

func (p Python) String() string {
	return "python"
}

func (p Python) GenerateSdkRecipe(providerName, path, outputPath string) []string {
	return BaseGenerateSdkCommand(providerName, path, outputPath, p.String())
}

func (p Python) CompileSdkRecipe(outputPath string) []string {

	// Named individual commands for ease of comprehension
	const (
		cdToPythonDir        = "cd {OutputPath}/python"
		cleanBinCmd          = "rm -rf ./bin/ ../python.bin/"
		copySrcUpOneLevelCmd = "cp -R . ../python.bin"
		moveSrcDirCmd        = "mv ../python.bin ./bin"
		deleteBinGoModCmd    = "rm -f ./bin/go.mod"
		createVenvCmd        = "python3 -m venv venv"
		installBuildDepCmd   = "./venv/bin/python -m pip install build==1.2.1"
		cdToBinDir           = "cd ./bin"
		buildPythonCmd       = "../venv/bin/python -m build ."
	)

	var compilePythonRecipie = []string{
		cdToPythonDir,
		cleanBinCmd,
		copySrcUpOneLevelCmd,
		moveSrcDirCmd,
		deleteBinGoModCmd,
		createVenvCmd,
		installBuildDepCmd,
		cdToBinDir,
		buildPythonCmd,
	}

	compilePythonCmd := strings.Join(compilePythonRecipie, joinCmdLineEnding)
	compilePythonCmd = strings.ReplaceAll(compilePythonCmd, "{OutputPath}", outputPath)
	return []string{compilePythonCmd}
}

func (p Python) PackageSdkRecipie() []string {
	return []string{}
}
