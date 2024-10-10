package terminal

import (
	"fmt"
	"os/exec"
	"strings"
)

func Execute(command string) (string, error) {
	args := strings.Split(command, " ")
	c := exec.Command(args[0], args[1:]...)

	output, err := c.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing command: %v\n", err.Error()), err

	}

	return fmt.Sprintf("Command output: %s\n", output), nil
}
