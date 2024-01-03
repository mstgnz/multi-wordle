package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
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
	fileName := fmt.Sprintf("./lang/%s/%d_letter_words.txt", lang, length)

	// Open the file
	file, err := os.Open(fileName)
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
	if err := scanner.Err(); err != nil {
		return 0, err
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

func ExistsLetter(word string, letter byte) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == letter {
			return true
		}
	}
	return false
}
