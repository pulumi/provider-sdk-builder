package lang

import (
	"testing"
)

func TestGenerateCopyCommand(t *testing.T) {
	tests := []struct {
		name       string
		sourcePath string
		sourceFile string
		outputPath string
		language   string
		destFile   string
		expected   string
	}{
		{
			name:       "README copy",
			sourcePath: "/test/provider",
			sourceFile: "README.md",
			outputPath: "/test/output",
			language:   "python",
			destFile:   "README.md",
			expected:   "cp -f \"/test/provider/README.md\" \"/test/output/python/README.md\"",
		},
		{
			name:       "LICENSE copy",
			sourcePath: "/test/provider",
			sourceFile: "LICENSE",
			outputPath: "/test/output",
			language:   "go",
			destFile:   "LICENSE",
			expected:   "cp -f \"/test/provider/LICENSE\" \"/test/output/go/LICENSE\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateCopyCommand(tt.sourcePath, tt.sourceFile, tt.outputPath, tt.language, tt.destFile)
			if result != tt.expected {
				t.Errorf("expected: %q\ngot:      %q", tt.expected, result)
			}
		})
	}
}