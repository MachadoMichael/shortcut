# Shortcut App Documentation

## Overview

**Shortcut** is a simple yet powerful command-line application designed to create aliases for long commands, particularly useful for SSH connections to various virtual machines (VMs). This tool was developed to alleviate the hassle of remembering user credentials and IP addresses for multiple VMs at work.

## Features

- **Create Aliases**: Easily save long commands as simple aliases.
- **Run Commands**: Quickly execute commands using their corresponding aliases.
- **Flexible Usage**: While designed for SSH connections, it can be used for any command.

## Installation

Ensure you have Go version **1.22.2** or later installed. You can download it from the official Go website.

## Usage

The application accepts two primary arguments: `save` and `run`.

### Saving an Alias

To save a new alias, use the following command:

```
go run main.go save <alias> "<command>"
```

<alias>: The name you want to assign to the command.
<command>: The long command you wish to alias.

Running a Command
To run a command associated with an alias, use:

```
go run main.go run <alias>
```

<alias>: The alias you previously saved.

Example
To save an SSH connection command:

```
go run main.go save myvm "ssh user@192.168.1.10"
```

To execute the saved command:

```
go run main.go run myvm
```

Code Explanation
The application uses a map structure to store aliases and their corresponding commands:

```
var commands = make(map[string]string)
```

Saving an Alias
When the save argument is provided, the application takes the alias and command, storing them in the map:

```
commands[alias] = command
```

Running a Command
When the run argument is invoked, the application retrieves and executes the command associated with the provided alias:

```
if cmd, exists := commands[alias]; exists {
    exec.Command(cmd).Run()
}
```

Conclusion
Shortcut simplifies the process of managing long commands by allowing users to create easy-to-remember aliases. This utility is especially beneficial for professionals who frequently connect to multiple virtual machines or execute lengthy commands in their workflows.
