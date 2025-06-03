package cmd

import "os/exec"

func Exec(args ...string) ([]byte, error) {

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, err
}
