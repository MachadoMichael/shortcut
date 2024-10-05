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

```bash
go run main.go save <alias> "<command>"
