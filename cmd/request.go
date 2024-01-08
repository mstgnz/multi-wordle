package main

type Request struct {
	Type     string   `json:"type"`
	Message  string   `json:"message"`
	Token    string   `json:"token"`
	Position Position `json:"position"`
}
