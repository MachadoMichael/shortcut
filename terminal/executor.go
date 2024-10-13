package terminal

import (
	"os"
	"os/exec"
)

// ExecuteInteractive runs a command interactively, attaching it to the terminal.
func ExecuteInteractive(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	// Attach the command's input/output to the current terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to complete
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return command, nil
}
