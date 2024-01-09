package main

import (
	"bufio"
	crtpto "crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

func HandleLog(message string, err error) {
	if err != nil {
		log.Println(message, err.Error())
	} else {
		log.Println(message)
	}
}

func ContainsValueInSlice(val any, array any) bool {
	arr := reflect.ValueOf(array)
	if arr.Kind() != reflect.Slice {
		return false
	}
	for i := 0; i < arr.Len(); i++ {
		if reflect.DeepEqual(val, arr.Index(i).Interface()) {
			return true
		}
	}
	return false
}

func SetAlphabet(lang string) []Alphabet {
	var alphabets []Alphabet
	var alphabet string

	switch lang {
	case "tr":
		alphabet = TurkishAlphabet
	default:
		alphabet = EnglishAlphabet
	}

	for _, letter := range alphabet {
		alphabets = append(alphabets, Alphabet{
			Letter: letter,
		})
	}
	return alphabets
}

// GetWords get random word
func GetWords(lang string, length int) (string, error) {
	// Specify the file name and path
	_, currentFile, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(currentFile), "Lang", lang, fmt.Sprintf("%d_letter_words.txt", length))

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	// Get the number of lines in the file
	lineCount, err := GetLineCount(file)
	if err != nil {
		return "", err
	}

	// Use a new source with the current time as the seed
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random line index
	randomLine := random.Intn(lineCount) + 1

	// Get the word from the selected line
	selectedWord, err := GetWordFromFile(file, randomLine)
	if err != nil {
		return "", err
	}

	return selectedWord, nil
}

// GetLineCount Get the number of lines in the file
func GetLineCount(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	// back to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		return lineCount, err
	}
	if err := scanner.Err(); err != nil {
		return lineCount, err
	}
	return lineCount, nil
}

// GetWordFromFile Get the word from the specified line
func GetWordFromFile(file *os.File, lineNo int) (string, error) {
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		if currentLine == lineNo {
			return scanner.Text(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New(fmt.Sprintf("Specified line not found: %d", lineNo))
}

// ExistsLetter check if there are letters in the word.
func ExistsLetter(word string, letter byte) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == letter {
			return true
		}
	}
	return false
}

// RandomName returns a random string of letters of Length n, using characters specified in randomStringSource.
func RandomName(n int) string {
	// randomStringSource is the source for generating random strings.
	const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321_"
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := crtpto.Prime(crtpto.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

// RandomColor generates a random color sequence for player icons.
func RandomColor() string {
	return fmt.Sprintf("rgb(%v,%v,%v)", RGB(), RGB(), RGB())
}

// RGB generates a random color.
func RGB() int {
	return rand.Intn(256)
}

// ContainsWord the predicted word is checked to see if it exists in the language used.
func ContainsWord(length int, lang, word string) (bool, error) {
	// Specify the file name and path
	_, currentFile, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(currentFile), "Lang", lang, fmt.Sprintf("%d_letter_words.txt", length))

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = file.Close()
	}()

	// scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true, nil
		}
	}
	return false, nil
}

// GenerateToken If the user is the same user when the user refreshes the page,
// it is necessary to take action accordingly.
// For this, we created a special token for each user.
func GenerateToken() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// InitRoom When the user first enters the page, they are greeted with a form that asks for the language of the game,
// number of sets, word length and number of attempts. This information can be entered or start with default values.
// if there is a value, it is set.
func InitRoom(request Request) (string, int, int, int) {
	var lang, limit, length, trial = "en", 3, 5, 5
	if request.Init.Lang == "tr" {
		lang = "tr"
	}
	if request.Init.Limit >= 1 && request.Init.Limit <= 9 {
		limit = request.Init.Limit
	}
	if request.Init.Length >= 1 && request.Init.Length <= 9 {
		length = request.Init.Length
	}
	if request.Init.Trial >= 1 && request.Init.Trial <= 9 {
		trial = request.Init.Trial
	}
	return lang, limit, length, trial
}

// FindTokenPlayerAndRoom when the user refreshes the page, the user and the room belonging to that user
// will be found with the token. if there is a player matching the token, the existing player will return,
// if not, a new login will be created.
func FindTokenPlayerAndRoom(request Request) (*Room, *Player) {
	var room *Room
	var player *Player
	player = PLAYERS.FindPlayerWithToken(request.Token)
	if player != nil {
		room = ROOMS.FindRoomWithWs(player.Conn)
	}
	return room, player
}
