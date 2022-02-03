package main

type Guesser struct {
	chars [5]byte
}

func NewGueser(word string) *Guesser {
	g := &Guesser{}
	for i := range word {
		g.chars[i] = word[i]
	}
	return g
}

func (g *Guesser) Guess(word string) (result guessResult) {
	charMap := make(map[byte]int, 5)
	for _, char := range g.chars {
		charMap[char]++
	}
	for i := range word {
		char := word[i]
		if _, found := charMap[char]; !found {
			continue
		}
		charMap[char]--
		if charMap[char] == 0 {
			delete(charMap, char)
		}
		if g.chars[i] == char {
			result[i] = DirectHit
		} else {
			result[i] = Hit
		}
	}
	return
}
