package main

type Request struct {
	Word string `json:"word"`
}

type Response struct {
	Status  bool   `json:"status"`
	Score   int    `json:"score"`
	Message string `json:"message"`
}
