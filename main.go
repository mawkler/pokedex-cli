package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mawkler/pokedex-cli/internal/cache"
	"github.com/mawkler/pokedex-cli/internal/cli"
	"github.com/mawkler/pokedex-cli/internal/cli/commands"
	"github.com/mawkler/pokedex-cli/internal/pokeapi"
	"github.com/mawkler/pokedex-cli/internal/pokedex"
)

func evaluate(input string, cfg *cli.Config, cliCommands map[string]commands.Command) error {
	cmd, args := cli.SplitInput(input)

	command, ok := cliCommands[cmd]
	if !ok {
		return fmt.Errorf("command not found: %s", input)
	}

	if err := command.Run(cfg, args...); err != nil {
		return err
	}

	return nil
}

func repl(scanner *bufio.Scanner, cfg cli.Config, cliCommands map[string]commands.Command) {
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

		err := evaluate(input, &cfg, cliCommands)
		if err != nil {
			println(err.Error())
		}
	}
}

func main() {
	cache := cache.NewCache(time.Minute * 2)
	pokeApiUrl := "https://pokeapi.co/api/v2"
	client := pokeapi.NewClient(pokeApiUrl, *http.DefaultClient, cache)
	pokedex := pokedex.NewPokedex()
	cfg := cli.NewConfig(client, pokedex)

	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := commands.NewCLICommandMap()

	// If CLI arguments were passed in
	input := strings.Join(os.Args[1:], " ")
	if len(input) > 0 {
		if err := evaluate(input, &cfg, cliCommands); err != nil {
			println(err.Error())
		}
		os.Exit(0)
	}

	// If no CLI arguments were passed in, start the REPL
	repl(scanner, cfg, cliCommands)
}
