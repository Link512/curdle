package main

import (
	"fmt"
	"strings"
)

const (
	tries = 6
)

func printGuesses(g []guessResult) {
	for i, gr := range g {
		fmt.Printf("%d ", i+1)
		for _, r := range gr {
			fmt.Printf("%c", finalChars[r])
		}
		fmt.Println()
	}
}

func main() {
	finalWord := getWord()
	g := NewGueser(finalWord)
	history := newHistory()
	guesses := make([]guessResult, 0)

	var (
		mainLine  = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
		underLine = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
		errLine   string
	)

	for i := 0; i < tries; i++ {
		fmt.Printf("\033[2J")
		if errLine != "" {
			fmt.Println(errLine)
			errLine = ""
		}
		if i > 0 {
			fmt.Println(string(mainLine[:]))
			fmt.Println(string(underLine[:]))
			fmt.Println()
			fmt.Println("History: ", history)
		}
		fmt.Printf("[%d] Your guess: ", i+1)

		var word string
		fmt.Scanln(&word)
		word = strings.ToLower(word)
		err := checkWord(word)
		if err != nil {
			errLine = err.Error()
			i--
			continue
		}
		guessResult := g.Guess(word)
		guesses = append(guesses, guessResult)
		for i, result := range guessResult {
			char := word[i]
			history.Add(char)
			mainLine[2*i] = char
			underLine[2*i] = guessChars[result]
		}

		if guessResult.FullGuess() {
			fmt.Printf("\033[2J")
			fmt.Println(string(mainLine[:]))
			fmt.Println(string(underLine[:]))
			fmt.Printf("\nCONGRATS!\n")
			printGuesses(guesses)
			return
		}
	}
	fmt.Printf("\033[2J")
	fmt.Println(string(mainLine[:]))
	fmt.Println(string(underLine[:]))
	fmt.Println("GAME OVER!")
	fmt.Println("The word was", finalWord)
}
