package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mawkler/pokedex-cli/cli"
	"github.com/mawkler/pokedex-cli/cli/commands"
)

func main() {
	cfg := cli.NewConfig()
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := commands.NewCLICommandMap()

	println("Welcome to the Pokedex!")

	for {
		print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		command, ok := cliCommands[input]
		if !ok {
			fmt.Printf("Command not found: %s", input)
			continue
		}

		if err := command.Run(&cfg); err != nil {
			fmt.Println(err)
		}
	}
}
