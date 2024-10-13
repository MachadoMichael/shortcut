package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Execute(command string) (string, error) {
	args := strings.Split(command, " ")
	c := exec.Command(args[0], args[1:]...)
	// Attach the Go program’s input, output, and error to the command
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	// Run the command and handle errors
	if err := c.Run(); err != nil {
		// Handle error
		os.Exit(1)
	}

	output, err := c.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing command: %v\n", err.Error()), err

	}

	return fmt.Sprintf("Command output: %s\n", output), nil
}

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

	return "Command executed successfully", nil
}
