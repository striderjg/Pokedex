package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"net/http"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next  string
	Prev  string
	cache *pokecache.Cache
}

var cmds map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	initConfig(&cfg)
	initCommands(&cfg)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmdWords := cleanInput(scanner.Text())
		if len(cmdWords) == 0 {
			continue
		}
		f, ok := cmds[cmdWords[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		f.callback(&cfg)
	}
}

// ==================== Command Handlers ===========

func commandMap(c *config) error {

	return _map(c, c.Next)
}

func commandMapb(c *config) error {
	if len(c.Prev) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	return _map(c, c.Prev)
}

// -- internal map that takes string url for which way we're going
func _map(c *config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("map failed to retrieve resource: %w", err)
	}
	defer res.Body.Close()

	var resJson mapJSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resJson); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	c.Next = resJson.Next
	c.Prev = resJson.Previous
	for _, n := range resJson.Results {
		fmt.Println(n.Name)
	}
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	helpText := "Welcome to the Pokedex!\nUsage:\n\n"
	for key, val := range cmds {
		helpText += key
		helpText += ": "
		helpText += val.description + "\n"
	}

	fmt.Print(helpText)
	return nil
}

// ======================== Utility Funcs

func cleanInput(text string) []string {
	wordList := strings.Fields(text)
	for i, word := range wordList {
		wordList[i] = strings.ToLower(word)
	}

	return wordList
}

func initConfig(c *config) {
	c.Next = "https://pokeapi.co/api/v2/location-area/"
	c.cache = pokecache.NewCache(5 * time.Second)

}

func initCommands(c *config) {
	cmds = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display consecutive locations every call",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
	}
}
