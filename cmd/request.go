package main

type Request struct {
	Type   string `json:"type"`
	Word   string `json:"word"`
	Player Player `json:"player"`
}
