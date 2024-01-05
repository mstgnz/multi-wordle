package main

type Request struct {
	Type     string   `json:"type"`
	Message  string   `json:"message"`
	Position Position `json:"position"`
}
