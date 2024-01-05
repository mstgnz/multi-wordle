package main

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Room    *Room  `json:"room"`
	// users belonging to the room
	Player  *Player   `json:"player"`
	Players []*Player `json:"players"`
}
