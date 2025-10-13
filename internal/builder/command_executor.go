package builder

import (
	"fmt"
	"io"
	"os/exec"
)

func ExecuteCommandSequence(commands []string, verbose bool, writer io.Writer) error {
	for index, command := range commands {
		if verbose {
			fmt.Fprintf(writer, "=== Command %d ===\n", index+1)
			fmt.Fprintf(writer, "%s\n", command)
			fmt.Fprintf(writer, "=== Command %d ===\n", index+1)
			fmt.Fprintf(writer, "\n")
		}

		// Execute the command using shell
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Stdout = writer
		cmd.Stderr = writer

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("%q FAILED: %w", command, err)
		}

		if verbose {
			fmt.Fprintf(writer, "\n")
		}
	}
	return nil
}
