package main

type Guesser struct {
	chars   [5]byte
	charMap map[byte]int
}

func NewGueser(word string) *Guesser {
	g := &Guesser{
		charMap: map[byte]int{},
	}
	for i := range word {
		g.charMap[word[i]]++
		g.chars[i] = word[i]
	}
	return g
}

func (g *Guesser) Guess(word string) (result guessResult) {
	for i := range word {
		char := word[i]
		if g.chars[i] == char {
			result[i] = DirectHit
			g.charMap[char]--
			if g.charMap[char] == 0 {
				delete(g.charMap, char)
			}
			continue
		}
		if _, found := g.charMap[char]; found {
			result[i] = Hit
			continue
		}
	}
	return
}
