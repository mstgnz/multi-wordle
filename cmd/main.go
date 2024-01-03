package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

var ROOMS []*Room
var WORDS []string

const (
	EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TurkishAlphabet = "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"
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
