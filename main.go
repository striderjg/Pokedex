package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmdWords := cleanInput(scanner.Text())
		if len(cmdWords) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", cmdWords[0])
	}

}

func cleanInput(text string) []string {
	wordList := strings.Fields(text)
	for i, word := range wordList {
		wordList[i] = strings.ToLower(word)
	}

	return wordList
}
