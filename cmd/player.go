package main

import (
	"golang.org/x/net/websocket"
)

// Player holds the player's information
type Player struct {
	Name       string          `json:"name"`
	Score      int             `json:"score"`
	IsGuessing bool            `json:"is_guessing"`
	Color      string          `json:"color"`
	Token      string          `json:"token"`
	Position   Position        `json:"position"`
	Conn       *websocket.Conn `json:"-"`
}

func NewPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		Name:     RandomName(5),
		Score:    0,
		Color:    RandomColor(),
		Token:    GenerateToken(),
		Position: Position{X: 100, Y: 100},
		Conn:     conn,
	}
	return PLAYERS.AddPlayer(player)
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

// SetIsGuessing is "true" if it is the guesser's turn.
func (p *Player) SetIsGuessing(isGuessing bool) {
	p.IsGuessing = isGuessing
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
