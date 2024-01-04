package main

import (
	"golang.org/x/net/websocket"
)

type Player struct {
	Name  string
	Score int
	Conn  *websocket.Conn
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

func NewPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		Name:  RandomName(5),
		Score: 0,
		Conn:  conn,
	}

	PLAYERS = append(PLAYERS, player)
	return player
}
