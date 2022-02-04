package main

import (
	"fmt"
	"strings"
)

const (
	tries   = 6
	wordLen = 5
)

type userGuess struct {
	word   string
	result guessResult
}

func printFinalGuesses(g []userGuess) {
	for i, gr := range g {
		fmt.Printf("%d ", i+1)
		for _, r := range gr.result {
			fmt.Printf("%c", finalChars[r])
		}
		fmt.Println()
	}
}

func printPreviousGuesses(g []userGuess) {
	var (
		mainLine  = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
		underLine = [10]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	)
	for _, gr := range g {
		for i, result := range gr.result {
			mainLine[2*i] = gr.word[i]
			underLine[2*i] = guessChars[result]
		}
		fmt.Println(string(mainLine[:]))
		fmt.Println(string(underLine[:]))
		fmt.Println()
	}
}

func printWin(guesses []userGuess) {
	clear()
	printPreviousGuesses(guesses)
	fmt.Printf("\nCONGRATS!\n\n")

	printFinalGuesses(guesses)
}

func printLoss(guesses []userGuess, word string) {
	clear()
	printPreviousGuesses(guesses)
	fmt.Println("GAME OVER!")

	fmt.Println("The word was", word)
	fmt.Println()

	printFinalGuesses(guesses)
}

func clear() {
	fmt.Printf("\033[2J")
}

func main() {
	finalWord := getWord()
	g := NewGueser(finalWord)
	history := newHistory()
	guesses := make([]userGuess, 0)

	for i := 0; i < tries; i++ {
		clear()
		if i > 0 {
			printPreviousGuesses(guesses)
			fmt.Println()
			fmt.Println("History: ", history)
		}

		var word string
		for {
			fmt.Printf("Guess #%d: ", i+1)
			fmt.Scanln(&word)
			word = strings.ToLower(word)
			err := validateWord(word)
			if err == nil {
				break
			}
			fmt.Println(err)
			word = ""
		}

		guessResult := g.Guess(word)
		guesses = append(guesses, userGuess{
			word:   word,
			result: guessResult,
		})
		for i := range guessResult {
			history.Add(word[i])
		}
		if guessResult.FullGuess() {
			printWin(guesses)
			return
		}
	}
	printLoss(guesses, finalWord)
}
