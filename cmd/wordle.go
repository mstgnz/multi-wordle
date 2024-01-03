package main

type Wordle struct {
	// the word to be predicted
	Word string
	// detail according to estimates
	Alphabet []Alphabet
}

func (w *Wordle) CheckWord(word string) {
	for i := 0; i < len(w.Word); i++ {
		if w.Word[i] == word[i] {
			// If the letter is present and in the correct position, it is green
			w.SetAlphabet(rune(w.Word[i]), "#00FF00")
		} else if ExistsLetter(word, w.Word[i]) {
			// If the letter is present but in the wrong position, it is yellow
			w.SetAlphabet(rune(w.Word[i]), "#FFFF00")
		} else {
			// If the letter is not present, it is gray
			w.SetAlphabet(rune(w.Word[i]), "#808080")
		}
	}
}

func (w *Wordle) SetAlphabet(letter rune, color string) {
	for _, v := range w.Alphabet {
		if v.Letter == letter {
			v.SetColor(color)
		}
	}
}

// Alphabet with each prediction the alphabet will be updated and shown to the players
type Alphabet struct {
	// each letter in the word
	Letter rune
	// If the letter is present and in the correct position, it is green
	// If the letter is present but in the wrong position, it is yellow
	// If the letter is not present, it is gray
	Color string
}

func (a *Alphabet) SetColor(color string) {
	a.Color = color
}
