package main

import (
	"golang.org/x/net/websocket"
)

// Players On each user connection, the user is assigned to the room. However, we keep a list of users (pointers) for easy user processing
type Players []*Player

// AddPlayer add player
func (p *Players) AddPlayer(player *Player) *Player {
	*p = append(*p, player)
	return player
}

// FindPlayer find player with ws
func (p *Players) FindPlayer(conn *websocket.Conn) *Player {
	for _, player := range *p {
		if player.Conn == conn {
			return player
		}
	}
	return nil
}

// FindToken find player with token
func (p *Players) FindToken(token string) *Player {
	for _, player := range *p {
		if player.Token == token {
			return player
		}
	}
	return nil
}

// DelPlayer delete player
func (p *Players) DelPlayer(conn *websocket.Conn) {
	for i, player := range *p {
		if player.Conn == conn {
			*p = append((*p)[:i], (*p)[i+1:]...)
			return
		}
	}
}
