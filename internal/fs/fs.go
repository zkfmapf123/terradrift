package fs

import (
	"fmt"
	"os/exec"
	"strings"
)

// ExecuteCommand executes a shell command and returns the output
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

// ExecuteShellCommand executes a shell command using /bin/sh
func ExecuteShellCommand(shellCommand string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", shellCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("shell command execution failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

// ExecuteCommandWithEnv executes a command with environment variables
func ExecuteCommandWithEnv(command string, env []string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution with env failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

// ExecuteCommandWithWorkingDir executes a command in a specific working directory
func ExecuteCommandWithWorkingDir(command string, workingDir string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workingDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution in directory failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}
