package main

import (
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

// ROOMS We keep a (pointer) list of all the rooms because the game is based on rooms
var ROOMS Rooms

// PLAYERS On each user connection, the user is assigned to the room. However, we keep a list of users (pointers) for easy user processing
// Since changes will be made over pointer addresses, the user in the room will also be updated.
var PLAYERS Players

const (
	EnglishAlphabet = "QWERTYUIOPASDFGHJKLZXCVBNM"
	TurkishAlphabet = "ERTYUIOPĞÜASDFGHJKLŞİZCVBNMÖÇ"
	MaxConnection   = 48
	Port            = 3000
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("public")))

	socket := NewSocket()
	http.Handle("/ws", websocket.Handler(socket.Handler))

	log.Println("Started Websocket on :", Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(Port), nil); err != nil {
		log.Fatal("HTTP server error:", err)
	}

}
