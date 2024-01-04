package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

// ROOMS We keep a (pointer) list of all the rooms because the game is based on rooms
var ROOMS Rooms

// PLAYERS On each user connection, the user is assigned to the room. However, we keep a list of users (pointers) for easy user processing
// Since changes will be made over pointer addresses, the user in the room will also be updated.
var PLAYERS Players

const (
	EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TurkishAlphabet = "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"
	MaxConnection   = 48
)

func main() {
	port := "3000"

	http.Handle("/", http.FileServer(http.Dir("public")))

	socket := NewSocket()
	http.Handle("/ws", websocket.Handler(socket.Handler))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Web socket failed to start:", err)
	}
}
