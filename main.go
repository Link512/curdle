package main

import (
	"fmt"
	"strings"
)

const (
	tries = 1
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

func printWin(mainLine, underLine [10]byte, guesses []guessResult) {
	fmt.Printf("\033[2J")
	fmt.Println(string(mainLine[:]))
	fmt.Println(string(underLine[:]))
	fmt.Printf("\nCONGRATS!\n\n")
	printGuesses(guesses)
}

func printLoss(mainLine, underLine [10]byte, guesses []guessResult, word string) {
	fmt.Printf("\033[2J")
	fmt.Println(string(mainLine[:]))
	fmt.Println(string(underLine[:]))
	fmt.Println("GAME OVER!")
	fmt.Println("The word was", word)
	fmt.Println()
	printGuesses(guesses)
}

func main() {
	finalWord := getWord()
	g := NewGueser(finalWord)
	history := newHistory()
	guesses := make([]guessResult, 0)

	var (
		mainLine  = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
		underLine = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	)

	for i := 0; i < tries; i++ {
		fmt.Printf("\033[2J")
		if i > 0 {
			fmt.Println(string(mainLine[:]))
			fmt.Println(string(underLine[:]))
			fmt.Println()
			fmt.Println("History: ", history)
		}
		fmt.Printf("Guess #%d: ", i+1)

		var word string
		fmt.Scanln(&word)
		word = strings.ToLower(word)
		err := checkWord(word)
		for err != nil {
			word = ""
			fmt.Println(err)
			fmt.Printf("Guess #%d: ", i+1)
			fmt.Scanln(&word)
			word = strings.ToLower(word)
			err = checkWord(word)
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
			printWin(mainLine, underLine, guesses)
			return
		}
	}
	printLoss(mainLine, underLine, guesses, finalWord)
}
