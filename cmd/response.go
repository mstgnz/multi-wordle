package main

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Room    *Room  `json:"room"`
}
