package main

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

type Room struct {
	ID           string
	Length       int
	Wordle       Wordle
	Players      map[*websocket.Conn]*Player
	PlayerTurn   *websocket.Conn
	PlayerTurnMu sync.Mutex
}

func NewRoom(lang string, length int) *Room {

	for _, room := range ROOMS {
		if len(room.Players) < 2 {
			return room
		}
	}

	getWord, err := GetWords(lang, length)
	if err != nil {
		getWord = "FAILED"
	}

	newRoom := &Room{
		ID:      fmt.Sprintf("room_%d", len(ROOMS)+1),
		Players: make(map[*websocket.Conn]*Player),
		Length:  5,
		Wordle: Wordle{
			Word:     getWord,
			Alphabet: SetAlphabet("en"),
		},
	}

	ROOMS = append(ROOMS, newRoom)
	return newRoom
}
