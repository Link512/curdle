package main

type dictionaryError struct {
	msg string
}

func (e *dictionaryError) Error() string {
	return string(e.msg)
}

var (
	errNotFiveChars = &dictionaryError{"input should be exactly 5 characters"}
)

func checkWord(s string) *dictionaryError {
	if len(s) != 5 {
		return errNotFiveChars
	}
	//TODO: check if actual word
	return nil
}

func getWord() string {
	return "shard"
}
