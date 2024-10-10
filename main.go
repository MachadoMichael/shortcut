package main

import (
	"fmt"
	"os"

	"github.com/MachadoMichael/shortcut/mapper"
	"github.com/MachadoMichael/shortcut/terminal"
	tui "github.com/MachadoMichael/shortcut/tui/fancy_list"
)

func main() {

	err := mapper.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numberOfArgs := len(os.Args)

	if numberOfArgs == 1 {
		tui.Init(mapper.CommandMapper.GetDictionary())
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

			mapper.CommandMapper.InsertInJson(alias, command)
			fmt.Println("The command has been saved")
			return
		}
	}

	if numberOfArgs == 3 {
		if os.Args[1] == "run" {
			alias := os.Args[2]
			command, err := mapper.CommandMapper.GetCommand(alias)
			if err != nil {
				fmt.Println("The command does not exist")
				return
			}

			output, err := terminal.Execute(command)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(output)
			return
		}
	}
}
