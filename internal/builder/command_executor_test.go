package builder

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecuteCommandSequence(t *testing.T) {
	tests := []struct {
		name        string
		commands    []string
		expectError bool
		checkOutput func(output string) bool
	}{
		{
			name:        "empty sequence",
			commands:    []string{},
			expectError: false,
			checkOutput: func(output string) bool {
				return output == ""
			},
		},
		{
			name:        "single echo command",
			commands:    []string{"echo 'hello world'"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "hello world") &&
					strings.Contains(output, "=== Command 1 ===")
			},
		},
		{
			name:        "multiple echo commands",
			commands:    []string{"echo 'first'", "echo 'second'"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "first") &&
					strings.Contains(output, "second") &&
					strings.Contains(output, "=== Command 1 ===") &&
					strings.Contains(output, "=== Command 2 ===")
			},
		},
		{
			name:        "command with arguments",
			commands:    []string{"pwd"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "/") && // Should contain path
					strings.Contains(output, "=== Command 1 ===")
			},
		},
		{
			name:        "invalid command",
			commands:    []string{"nonexistentcommand123"},
			expectError: true,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "=== Command 1 ===")
			},
		},
		{
			name:        "mixed valid and invalid commands",
			commands:    []string{"echo 'before error'", "nonexistentcommand123"},
			expectError: true,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "before error") &&
					strings.Contains(output, "=== Command 1 ===") &&
					strings.Contains(output, "=== Command 2 ===")
			},
		},
		{
			name:        "skip empty commands",
			commands:    []string{"echo 'test'", "", "echo 'after empty'"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Contains(output, "test") &&
					strings.Contains(output, "after empty") &&
					strings.Contains(output, "=== Command 1 ===") &&
					strings.Contains(output, "=== Command 3 ===")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Execute the sequence
			err := ExecuteCommandSequence(tt.commands, true, &buf)

			// Check error expectation
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Check output
			output := buf.String()
			if !tt.checkOutput(output) {
				t.Errorf("output check failed. Got output: %q", output)
			}
		})
	}
}

func TestExecuteCommandSequenceSingleCommand(t *testing.T) {
	commands := []string{"echo 'test output'"}

	var buf bytes.Buffer
	err := ExecuteCommandSequence(commands, true, &buf)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "test output") {
		t.Errorf("expected output to contain 'test output', got: %q", output)
	}

	if !strings.Contains(output, "=== Command 1 ===") {
		t.Errorf("expected output to contain command header, got: %q", output)
	}
}

func TestExecuteCommandSequenceErrorHandling(t *testing.T) {
	commands := []string{
		"echo 'before error'",
		"exit 1",             // This will cause an error
		"echo 'after error'", // This should not execute
	}

	var buf bytes.Buffer
	err := ExecuteCommandSequence(commands, true, &buf)

	if err == nil {
		t.Error("expected error but got none")
	}

	output := buf.String()

	// Should contain output from the first command
	if !strings.Contains(output, "before error") {
		t.Errorf("expected output to contain 'before error', got: %q", output)
	}

	// Should NOT contain output from the third command
	if strings.Contains(output, "after error") {
		t.Errorf("output should not contain 'after error' since execution stopped on error, got: %q", output)
	}

	// Should contain error details
	if !strings.Contains(err.Error(), "\"exit 1\" FAILED:") {
		t.Errorf("expected error message to contain 'command FAILED', got: %v", err)
	}
}

func TestExecuteCommandSequenceOutputFormatting(t *testing.T) {
	commands := []string{"echo 'line1'", "echo 'line2'"}

	var buf bytes.Buffer
	err := ExecuteCommandSequence(commands, true, &buf)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	output := buf.String()

	// Verify command headers are present
	if !strings.Contains(output, "=== Command 1 ===") {
		t.Errorf("expected first command header, got: %q", output)
	}

	if !strings.Contains(output, "=== Command 2 ===") {
		t.Errorf("expected second command header, got: %q", output)
	}

	// Verify both outputs are present
	if !strings.Contains(output, "line1") {
		t.Errorf("expected 'line1' in output, got: %q", output)
	}

	if !strings.Contains(output, "line2") {
		t.Errorf("expected 'line2' in output, got: %q", output)
	}
}
