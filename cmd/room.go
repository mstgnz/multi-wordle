package main

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

// Room rooms are for a maximum of two people. the game is based on two people.
// socket operations will take place through the room.
// room object will be sent as a response. room object is the room of the user taking action.
type Room struct {
	// Room Name
	Name string `json:"name"`
	// Limit how many matches will be played
	Lang string `json:"lang"`
	// Messages intra-room correspondence
	Limit int `json:"limit"`
	// Word Length
	Length int `json:"len"`
	// Trial Number of word prediction trials
	Trial int `json:"trial"`
	// Lang word language to be predicted
	Messages []string `json:"messages"`
	// Wordle It provides the word to be guessed and the necessary checks and coloring for each guess.
	Wordle *Wordle `json:"wordle"`
	// Matches Prediction information for all matches
	Matches []*Wordle `json:"matches"`
	// Players It holds the users in the room. 2 users.
	Players map[*websocket.Conn]*Player `json:"-"`
	// PlayerTurn Holds the user who will make a word prediction.
	PlayerTurn *websocket.Conn `json:"-"`
	// Mutex Locking mechanism for healthy word prediction.
	Mutex sync.Mutex `json:"-"`
}

// NewRoom When a user is connected, if there is a room with 1 user,
// the user will enter that room, if not,
// a new room will be created and the user will enter there.
func NewRoom(request Request) (*Room, error) {
	// initialized room settings
	lang, limit, length, trial := InitRoom(request)
	for _, room := range ROOMS {
		if room.Lang == lang && room.Limit == limit && room.Length == length && room.Trial == trial && len(room.Players) < 2 {
			return room, nil
		}
	}

	getWord, err := GetWords(lang, length)
	if err != nil {
		return nil, err
	}

	newRoom := &Room{
		Name:    fmt.Sprintf("room_%d", len(ROOMS)+1),
		Players: make(map[*websocket.Conn]*Player),
		Lang:    lang,
		Limit:   limit,
		Length:  length,
		Trial:   trial,
		Wordle: &Wordle{
			Word:      getWord,
			Forecasts: make([]Forecasts, 0),
			Alphabet:  SetAlphabet(lang),
		},
	}
	return ROOMS.AddRoom(newRoom), nil
}

// AddMessage append message
func (r *Room) AddMessage(message string) {
	r.Messages = append(r.Messages, message)
}

// GetPlayers get players from map object
func (r *Room) GetPlayers() []*Player {
	var players []*Player
	for _, player := range r.Players {
		if player != nil {
			players = append(players, player)
		}
	}
	return players
}

// Reset New match
func (r *Room) Reset() (*Room, error) {
	getWord, err := GetWords(r.Lang, r.Length)
	if err != nil {
		return nil, err
	}
	r.Matches = append(r.Matches, r.Wordle)
	r.Wordle = &Wordle{
		Word:      getWord,
		Forecasts: make([]Forecasts, 0),
		Alphabet:  SetAlphabet(r.Lang),
	}
	return r, nil
}

// CheckWord In each prediction attempt, the word to be predicted is compared with the predicted word.
// Coloring on the alphabet sequence according to the comparison.
func (r *Room) CheckWord(word string, player *Player) {
	var forecasts []Forecast
	score := 0
	for i := 0; i < len(r.Wordle.Word); i++ {
		if r.Wordle.Word[i] == word[i] {
			// If the letter is present and in the correct position, it is green
			r.Wordle.SetAlphabet(rune(word[i]), "#00FF00")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#00FF00"})
			// 5 points for finding the correct letter in the word.
			// If a player finds the letter in the word and the location is wrong and the next player sees the letter and locates it correctly, then 2 points from 5-3.
			score += 5
		} else if ExistsLetter(r.Wordle.Word, word[i]) {
			// If the letter is present but in the wrong position, it is yellow
			r.Wordle.SetAlphabet(rune(word[i]), "#FFFF00")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#FFFF00"})
			// 3 points for finding the letter in the word but misplacing it.
			score += 2
		} else {
			// If the letter is not present, it is gray
			r.Wordle.SetAlphabet(rune(word[i]), "#808080")
			forecasts = append(forecasts, Forecast{Letter: rune(word[i]), Color: "#808080"})
			// if there is no letter in the word, no penalty for the first use but -1 point for the second use.
			score -= 1
		}
	}
	player.Score += score
	r.Wordle.Forecasts = append(r.Wordle.Forecasts, Forecasts{
		Word:     word,
		Score:    score,
		Forecast: forecasts,
		Player:   player,
	})
}
