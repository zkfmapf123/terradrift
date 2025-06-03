package cmd

import (
	"fmt"
	"os/exec"
)

// ExecuteShell executes a shell command with given arguments
func ExecuteShell(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no command provided")
	}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v, output: %s", err, string(output))
	}
	return string(output), nil
}
