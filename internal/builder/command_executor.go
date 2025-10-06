package builder

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommandSequence(commands []string, verbose bool) (output string, err error) {
	var outputBuilder strings.Builder

	for index, command := range commands {
		if verbose {
			outputBuilder.WriteString(fmt.Sprintf("=== Command %d ===\n", index+1))
			outputBuilder.WriteString(command)
			outputBuilder.WriteString(fmt.Sprintf("\n===  Command %d ===\n", index+1))
			outputBuilder.WriteString("\n")
		}

		// Execute the command using shell
		cmd := exec.Command("/bin/sh", "-c", command)
		output, err := cmd.CombinedOutput()
		stringOutput := strings.TrimSpace(string(output))
		outputBuilder.WriteString(stringOutput)
		outputBuilder.WriteString("\n")

		if err != nil {
			return outputBuilder.String(), fmt.Errorf("%q FAILED: %w", command, err)
		}
		if verbose {
			outputBuilder.WriteString("\n")
		}
	}
	return outputBuilder.String(), nil
}
