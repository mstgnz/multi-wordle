package main

// Wordle Keeps an array of the word to be predicted and the alphabet of the selected language.
type Wordle struct {
	// Word the word to be predicted - json will not be in the output but attached for test purposes
	Word string `json:"word"`
	// Forecasts keeps a comparison of all attempts and words attempted
	Forecasts []Forecasts `json:"forecasts"`
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
			w.SetAlphabet(rune(word[i]), "#00FF00")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#00FF00"})
		} else if ExistsLetter(w.Word, word[i]) {
			// If the letter is present but in the wrong position, it is yellow
			w.SetAlphabet(rune(word[i]), "#FFFF00")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#FFFF00"})
		} else {
			// If the letter is not present, it is gray
			w.SetAlphabet(rune(word[i]), "#808080")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#808080"})
		}
	}
	w.Forecasts = append(w.Forecasts, Forecasts{
		Word:     word,
		Forecast: forecasts,
	})
}

// SetAlphabet Triggered by CheckWord and the corresponding rune is found in the Alphabet slice
func (w *Wordle) SetAlphabet(letter rune, color string) {
	for i := range w.Alphabet {
		if w.Alphabet[i].Letter == letter {
			w.Alphabet[i].SetColor(color)
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

// Forecasts keeps all predictions
type Forecasts struct {
	Word     string     `json:"word"`
	Forecast []Forecast `json:"forecast"`
}
