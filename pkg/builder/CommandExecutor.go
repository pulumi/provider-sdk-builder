package builder

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pulumi/provider-sdk-builder/pkg/shell"
)

func ExecuteCommandSequence(sequence shell.ShellCommandSequence) (output string, err error) {
	var outputBuilder strings.Builder
	var hasOutput bool

	for i, command := range sequence.Commands {
		addSeparator(&outputBuilder, hasOutput)
		addCommandHeader(&outputBuilder, i, command)

		// Execute the command using shell
		cmd := exec.Command("/bin/sh", "-c", command)
		cmdOutput, err := cmd.CombinedOutput()

		if err != nil {
			addErrorOutput(&outputBuilder, cmdOutput)
			return outputBuilder.String(), fmt.Errorf("command failed: %s, error: %w", command, err)
		}

		if addCommandOutput(&outputBuilder, cmdOutput) {
			hasOutput = true
		}
	}

	return outputBuilder.String(), nil
}

// addSeparator adds a separator line between commands if needed
func addSeparator(builder *strings.Builder, hasOutput bool) {
	if hasOutput {
		builder.WriteString("\n")
	}
}

// addCommandHeader adds a formatted command header to the output
func addCommandHeader(builder *strings.Builder, commandIndex int, command string) {
	builder.WriteString(fmt.Sprintf("=== Command %d: %s ===\n", commandIndex+1, command))
}

// addCommandOutput adds the command output to the builder, ensuring proper newline handling
func addCommandOutput(builder *strings.Builder, output []byte) bool {
	if len(output) == 0 {
		return false
	}

	builder.WriteString(string(output))
	if !strings.HasSuffix(string(output), "\n") {
		builder.WriteString("\n")
	}
	return true
}

// addErrorOutput adds command output when an error occurs
func addErrorOutput(builder *strings.Builder, output []byte) {
	if len(output) > 0 {
		builder.WriteString(string(output))
		builder.WriteString("\n")
	}
}
