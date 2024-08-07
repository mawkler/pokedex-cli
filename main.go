package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mawkler/pokedex-cli/internal/cache"
	"github.com/mawkler/pokedex-cli/internal/cli"
	"github.com/mawkler/pokedex-cli/internal/cli/commands"
)

func main() {
	cache := cache.NewCache(time.Millisecond * 10)
	cache.Add("key", []byte("value"))

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
