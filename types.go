package main

type guess uint8

const (
	Miss guess = iota
	Hit
	DirectHit
)

var (
	guessChars = [3]byte{
		' ',
		'*',
		'x',
	}
	finalChars = [3]byte{
		'_',
		'*',
		'x',
	}
)

type guessResult [5]guess

func (g guessResult) FullGuess() bool {
	for _, r := range g {
		if r != DirectHit {
			return false
		}
	}
	return true
}
