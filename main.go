package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mawkler/pokedex-cli/internal/cache"
	"github.com/mawkler/pokedex-cli/internal/cli"
	"github.com/mawkler/pokedex-cli/internal/cli/commands"
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
	cache := cache.NewCache(time.Millisecond * 10)
	cache.Add("key", []byte("value"))

	cfg := cli.NewConfig()
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
