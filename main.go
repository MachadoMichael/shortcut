package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.Open("dictionary.json")
	if err != nil {
		fmt.Println("read file")
		log.Fatal(err.Error())
	}
	defer file.Close()

	bytesValue, _ := io.ReadAll(file)
	if err != nil {
		fmt.Println("read bytes")
		log.Fatal(err.Error())
	}

	if len(bytesValue) == 0 {
		log.Fatal("The file is empty")
	}

	var dictionary map[string]string
	err = json.Unmarshal(bytesValue, &dictionary)
	if err != nil {
		fmt.Println("unmarshal")
		log.Fatal(err.Error())
	}

	numberOfArgs := len(os.Args)

	if numberOfArgs == 4 {
		if os.Args[1] == "save" {
			alias := os.Args[2]
			command := os.Args[3]

			insertInJson(alias, command, dictionary)
			fmt.Println("The command has been saved")
			return
		}
	}

	if numberOfArgs == 3 {
		if os.Args[1] == "run" {
			alias := os.Args[2]
			command := dictionary[alias]
			fmt.Println(command)
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

func insertInJson(alias, data string, dictionary map[string]string) {
	dictionary[alias] = data
	saveJson(dictionary)
}

func saveJson(dictionary map[string]string) {
	json_data, _ := json.Marshal(dictionary)
	err := os.WriteFile("./dictionary.json", json_data, 0644)
	if err != nil {
		panic(err)
	}
}
