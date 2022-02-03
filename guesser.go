package main

type Guesser struct {
	chars   [5]byte
	charMap map[byte]struct{}
}

func NewGueser(word string) *Guesser {
	g := &Guesser{
		charMap: map[byte]struct{}{},
	}
	for i := range word {
		g.charMap[word[i]] = struct{}{}
		g.chars[i] = word[i]
	}
	return g
}

func (g *Guesser) Guess(word string) (result guessResult) {
	for i := range word {
		char := word[i]
		if g.chars[i] == char {
			result[i] = DirectHit
			continue
		}
		if _, found := g.charMap[char]; found {
			result[i] = Hit
			continue
		}
	}
	return
}
