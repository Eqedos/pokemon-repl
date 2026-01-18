// Package main is the entry point for the Pokedex CLI application.
// It provides a REPL interface to explore Pokemon locations using the PokeAPI.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/eqedos/repl/internal/pokeapi"
)

// config holds the application state.
type config struct {
	client  *pokeapi.Client
	nextURL *string
	prevURL *string
	pokedex map[string]pokeapi.Pokemon
}

// cliCommand represents a command that can be executed in the Pokedex REPL.
type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func main() {
	// Initialize application state
	client := pokeapi.NewClient()
	firstURL := client.GetFirstLocationAreasURL()

	cfg := &config{
		client:  client,
		nextURL: &firstURL,
		prevURL: nil,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	// Start the REPL
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := cleanInput(input)

		if len(args) == 0 {
			continue
		}

		cmdName := args[0]
		cmd, exists := getCommands()[cmdName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		if err := cmd.callback(cfg, args[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

// getCommands returns all available CLI commands.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 Pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all Pokemon in a location (usage: explore <location-name>)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon (usage: catch <pokemon-name>)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View details of a caught Pokemon (usage: inspect <pokemon-name>)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all Pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

// cleanInput normalizes user input by splitting on whitespace and converting to lowercase.
func cleanInput(text string) []string {
	words := strings.Fields(text)
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = strings.ToLower(word)
	}
	return result
}

// commandHelp displays all available commands and their descriptions.
func commandHelp(cfg *config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range getCommands() {
		fmt.Printf("  %s: %s\n", name, cmd.description)
	}
	fmt.Println()
	return nil
}

// commandExit terminates the Pokedex application.
func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandMap displays the next 20 Pokemon location areas.
func commandMap(cfg *config, args []string) error {
	if cfg.nextURL == nil {
		fmt.Println("You're on the last page")
		return nil
	}

	resp, err := cfg.client.GetLocationAreas(*cfg.nextURL)
	if err != nil {
		return err
	}

	// Update pagination state
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	// Display locations
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// commandMapb displays the previous 20 Pokemon location areas.
func commandMapb(cfg *config, args []string) error {
	if cfg.prevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := cfg.client.GetLocationAreas(*cfg.prevURL)
	if err != nil {
		return err
	}

	// Update pagination state
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	// Display locations
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// commandExplore displays all Pokemon that can be encountered in a given location.
func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location name (e.g., 'explore canalave-city-area')")
	}

	locationName := args[0]

	resp, err := cfg.client.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Location.Name)
	fmt.Println("Found Pokemon:")

	if len(resp.PokemonEncounters) == 0 {
		fmt.Println("  No Pokemon found in this area.")
	} else {
		for _, encounter := range resp.PokemonEncounters {
			fmt.Printf("  - %s\n", encounter.Pokemon.Name)
		}
	}

	return nil
}

// commandCatch attempts to catch a Pokemon and add it to the user's Pokedex.
func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name (e.g., 'catch pikachu')")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Fetch Pokemon data
	pokemon, err := cfg.client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Calculate catch chance based on base experience
	// Higher base experience = harder to catch
	// Base experience ranges from ~36 (low) to ~608 (legendary)
	// We'll use a threshold approach: random number must exceed a scaled value
	const maxBaseExp = 400
	catchThreshold := pokemon.BaseExperience
	catchThreshold = min(catchThreshold, maxBaseExp)

	// Generate random number between 0 and maxBaseExp
	// If random >= catchThreshold, the Pokemon is caught
	roll := rand.Intn(maxBaseExp)

	if roll >= catchThreshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonName] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// commandInspect displays details of a caught Pokemon from the user's Pokedex.
func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name (e.g., 'inspect pikachu')")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

// commandPokedex lists all Pokemon the user has caught.
func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Try catching some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
