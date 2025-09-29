package builder

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommandSequence(commands []string) (output string, err error) {
	var outputBuilder strings.Builder

	for index, command := range commands {
		outputBuilder.WriteString(fmt.Sprintf("=== Command %d: %s ===\n", index+1, command))

		// Execute the command using shell
		cmd := exec.Command("/bin/sh", "-c", command)
		output, err := cmd.CombinedOutput()
		stringOutput := strings.TrimSpace(string(output))
		outputBuilder.WriteString(stringOutput)
		outputBuilder.WriteString("\n")

		if err != nil {
			return outputBuilder.String(), fmt.Errorf("%q failed: %w", command, err)
		}
	}
	return outputBuilder.String(), nil
}
