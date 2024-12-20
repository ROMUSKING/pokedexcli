package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandHandler(words []string) error {

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	if len(words) < 1 {
		return nil
	}
	command, ok := commands[words[0]]
	if !ok {
		return fmt.Errorf("command unknown")
	}
	return command.callback()
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

`,
	)
	for _, command := range commands {
		fmt.Printf(" %s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
