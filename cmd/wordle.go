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
	Score    int        `json:"score"`
	Forecast []Forecast `json:"forecast"`
	// guessing player
	Player *Player `json:"player"`
}
