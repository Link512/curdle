package main

import "testing"

func TestCheckWord(t *testing.T) {
	testCases := []struct {
		name          string
		word          string
		expectedError *dictionaryError
	}{
		{
			name:          "empty",
			expectedError: errNotFiveChars,
		},
		{
			name:          "under 5 chars",
			word:          "pie",
			expectedError: errNotFiveChars,
		},
		{
			name:          "over 5 chars",
			word:          "sprinkle",
			expectedError: errNotFiveChars,
		},
		{
			name:          "not a word",
			word:          "abcde",
			expectedError: errInvalidWord,
		},
		{
			name: "good value",
			word: getWord(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkWord(tc.word)
			if tc.expectedError != err {
				t.Errorf("actual error: %+v, expected error: %+v", err, tc.expectedError)
			}
		})
	}
}
