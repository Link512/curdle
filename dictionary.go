package main

import (
	"fmt"
	"math/rand"
	"regexp"
)

type dictionaryError struct {
	msg string
}

func (e *dictionaryError) Error() string {
	return string(e.msg)
}

var (
	errNotEnoughChars = &dictionaryError{fmt.Sprintf("not a %d letter word", wordLen)}
	errNotAlpha       = &dictionaryError{"word contains invalid characters"}
	errInvalidWord    = &dictionaryError{"word does not exist"}

	wordRegex = regexp.MustCompile(fmt.Sprintf(`^[a-z]{%d}$`, wordLen))
)

func validateWord(s string) *dictionaryError {
	if len(s) != wordLen {
		return errNotEnoughChars
	}
	if !wordRegex.MatchString(s) {
		return errNotAlpha
	}
	if _, found := allowedWords[s]; !found {
		return errInvalidWord
	}
	return nil
}

func getWord() string {
	return wordList[rand.Intn(len(wordList))]
}
