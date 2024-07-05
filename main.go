package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// Hangman project.

	words := []string{"php", "js", "golang"}
	word := words[rand.Intn(len(words))]
	lives := 5
	blanks := []string{}
	for range word {
		blanks = append(blanks, "_")
	}

	fmt.Printf("Word: %s Letter: ", strings.Join(blanks, ""))

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	for _, inputLetter := range input {
		correctGuess := false

		for i, wordLetter := range word {
			if inputLetter == wordLetter {
				blanks[i] = string(inputLetter)
				correctGuess = true
			}
		}
		if correctGuess {
			lives--
		}
	}

	if lives <= 0 {
		fmt.Printf("❤️  0, Word: %s - sorry, you lost!\n", word)
	}

	if word == strings.Join(blanks, "") {
		fmt.Printf("❤️ %d, Word: %s - you won, congrats!\n", lives, word)
	}
}
