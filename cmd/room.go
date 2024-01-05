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
	// Room ID
	ID string `json:"id"`
	// Word Length
	Length int `json:"len"`
	// Trial number of trials
	Trial int `json:"trial"`
	// Messages intra-room correspondence
	Messages []string `json:"messages"`
	// Wordle It provides the word to be guessed and the necessary checks and coloring for each guess.
	Wordle Wordle `json:"wordle"`
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
func NewRoom(lang string, length, trial int) (*Room, error) {

	for _, room := range ROOMS {
		if len(room.Players) < 2 {
			return room, nil
		}
	}

	getWord, err := GetWords(lang, length)
	if err != nil {
		return nil, err
	}

	newRoom := &Room{
		ID:      fmt.Sprintf("room_%d", len(ROOMS)+1),
		Players: make(map[*websocket.Conn]*Player),
		Length:  length,
		Trial:   trial,
		Wordle: Wordle{
			Word:     getWord,
			Alphabet: SetAlphabet(lang),
		},
	}

	ROOMS = append(ROOMS, newRoom)
	return newRoom, nil
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
