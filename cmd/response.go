package main

type Response struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Wordles []Wordle `json:"wordles"`
	Player  Player   `json:"player"`
	Players []Player `json:"players"`
}
