package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if input.Scan() {
			text := input.Text()
			words := cleanInput(text)
			err := commandHandler(words)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
