package main

import (
	"bufio"
	"log"
	"os"

	"pokedex-cli/cli"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := cli.NewCLICommandMap()

	println("Welcome to the Pokedex!")

	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		input := scanner.Text()
		command, ok := cliCommands[input]
		if !ok {
			log.Default().Fatalf("Command not found: %s", input)
		}

		command.Run()
	}
}
