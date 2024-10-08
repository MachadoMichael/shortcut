package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/MachadoMichael/shortcut/mapper"
	tui "github.com/MachadoMichael/shortcut/tui/fancy_list"
)

func main() {

	m := &mapper.Mapper{}

	dic, err := m.BuildMap("/Users/michael/Projects/shortcut/dictionary.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numberOfArgs := len(os.Args)

	if numberOfArgs == 1 {
		tui.Init(dic)
		return
	}

	if numberOfArgs != 3 && numberOfArgs != 4 {
		fmt.Println("Invalid number of arguments")
		return
	}

	if numberOfArgs == 4 {
		if os.Args[1] == "save" {
			alias := os.Args[2]
			command := os.Args[3]

			m.InsertInJson(alias, command, dic)
			fmt.Println("The command has been saved")
			return
		}
	}

	if numberOfArgs == 3 {
		if os.Args[1] == "run" {
			alias := os.Args[2]
			command, exist := dic[alias]
			if !exist {
				fmt.Println("The command does not exist")
				return
			}

			args := strings.Split(command, " ")
			c := exec.Command(args[0], args[1:]...)

			output, err := c.CombinedOutput()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}

			fmt.Printf("Command output: %s\n", output)
			return
		}
	}
}
