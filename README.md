# Pokedex CLI

A command-line Pokedex application built in Go that lets you explore Pokemon locations, catch Pokemon, and build your collection using data from the [PokeAPI](https://pokeapi.co/).

## Features

- Browse Pokemon location areas with pagination
- Explore locations to discover which Pokemon can be found there
- Catch Pokemon with a chance-based system (rarer Pokemon are harder to catch)
- Build your personal Pokedex collection
- Inspect caught Pokemon to view their stats and types
- Response caching to minimize API calls

## Installation

```bash
go build -o pokedex ./cmd/pokedex
```

## Usage

Run the application:

```bash
./pokedex
```

### Commands

| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `map` | List the next 20 Pokemon locations |
| `mapb` | List the previous 20 Pokemon locations |
| `explore <location>` | Show all Pokemon in a location |
| `catch <pokemon>` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | View details of a caught Pokemon |
| `pokedex` | List all Pokemon you have caught |
| `exit` | Exit the application |

### Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria...
Found Pokemon:
  - tentacool
  - tentacruel
  - magikarp
  - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Your Pokedex:
  - magikarp
```

## Project Structure

```
.
├── cmd/
│   └── pokedex/
│       ├── main.go         # Entry point, REPL, and commands
│       └── main_test.go    # Tests
├── internal/
│   ├── cache/
│   │   ├── cache.go        # Thread-safe cache with TTL
│   │   └── cache_test.go   # Cache tests
│   └── pokeapi/
│       ├── client.go       # API client with caching
│       └── types.go        # API response types
├── go.mod
└── README.md
```

## Running Tests

```bash
go test ./...
```

## License

MIT
