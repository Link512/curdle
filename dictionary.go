package main

import "math/rand"

type dictionaryError struct {
	msg string
}

func (e *dictionaryError) Error() string {
	return string(e.msg)
}

var (
	errNotFiveChars = &dictionaryError{"input should be exactly 5 characters"}
	errInvalidWord  = &dictionaryError{"word not in word list"}
)

func checkWord(s string) *dictionaryError {
	if len(s) != 5 {
		return errNotFiveChars
	}
	if _, found := allowedWords[s]; !found {
		return errInvalidWord
	}
	return nil
}

func getWord() string {
	return wordList[rand.Intn(len(wordList))]
}
