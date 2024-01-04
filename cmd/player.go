package main

import (
	"golang.org/x/net/websocket"
)

type Player struct {
	Name        string
	Score       int
	WordToGuess []string
	IsGuessing  bool
	Conn        *websocket.Conn
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

func (p *Player) SetIsGuessing(isGuessing bool) {
	p.IsGuessing = isGuessing
}

func (p *Player) AddWordToGuess(guess string) {
	p.WordToGuess = append(p.WordToGuess, guess)
}

func NewPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		Name:        RandomName(5),
		Score:       0,
		WordToGuess: nil,
		IsGuessing:  false,
		Conn:        conn,
	}

	PLAYERS = append(PLAYERS, player)
	return player
}
