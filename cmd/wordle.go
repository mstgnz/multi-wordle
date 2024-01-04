package main

// Wordle Keeps an array of the word to be predicted and the alphabet of the selected language.
type Wordle struct {
	// the word to be predicted
	Word string `json:"word"`
	// Forecasts keeps a comparison of all attempts and words attempted
	Forecasts map[string][]Forecast `json:"forecasts"`
	// detail according to estimates
	Alphabet []Alphabet `json:"alphabet"`
}

// CheckWord In each prediction attempt, the word to be predicted is compared with the predicted word.
// Coloring on the alphabet sequence according to the comparison.
func (w *Wordle) CheckWord(word string) {
	var forecasts []Forecast
	for i := 0; i < len(w.Word); i++ {
		if w.Word[i] == word[i] {
			// If the letter is present and in the correct position, it is green
			w.SetAlphabet(rune(w.Word[i]), "#00FF00")
			forecasts = append(forecasts, Forecast{Letter: rune(w.Word[i]), Color: "#00FF00"})
		} else if ExistsLetter(word, w.Word[i]) {
			// If the letter is present but in the wrong position, it is yellow
			w.SetAlphabet(rune(w.Word[i]), "#FFFF00")
			forecasts = append(forecasts, Forecast{Letter: rune(w.Word[i]), Color: "#FFFF00"})
		} else {
			// If the letter is not present, it is gray
			w.SetAlphabet(rune(w.Word[i]), "#808080")
			forecasts = append(forecasts, Forecast{Letter: rune(w.Word[i]), Color: "#808080"})
		}
	}
	w.Forecasts[word] = forecasts
}

// SetAlphabet Triggered by CheckWord and the corresponding rune is found in the Alphabet slice
func (w *Wordle) SetAlphabet(letter rune, color string) {
	for _, v := range w.Alphabet {
		if v.Letter == letter {
			v.SetColor(color)
		}
	}
}

// Alphabet Holds the alphabet of the selected language when the game is set.
type Alphabet struct {
	// each letter in the word
	Letter rune `json:"letter"`
	// If the letter is present and in the correct position, it is green
	// If the letter is present but in the wrong position, it is yellow
	// If the letter is not present, it is gray
	Color string `json:"color"`
}

// SetColor It is triggered by SetAlphabet and the color change is applied for the corresponding letter.
func (a *Alphabet) SetColor(color string) {
	a.Color = color
}

// Forecast holds the alphabet of predicted words.
type Forecast struct {
	// each letter in the word
	Letter rune `json:"letter"`
	// If the letter is present and in the correct position, it is green
	// If the letter is present but in the wrong position, it is yellow
	// If the letter is not present, it is gray
	Color string `json:"color"`
}

// SetColor It is triggered by SetAlphabet and the color change is applied for the corresponding letter.
func (f *Forecast) SetColor(color string) {
	f.Color = color
}
