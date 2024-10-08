package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeCommand(command string) error {
	args := strings.Split(command, " ")
	c := exec.Command(args[0], args[1:]...)

	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return err
	}

	fmt.Printf("Command output: %s\n", output)
	return nil
}
