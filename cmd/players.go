package main

import (
	"golang.org/x/net/websocket"
)

// Players On each user connection, the user is assigned to the room. However, we keep a list of users (pointers) for easy user processing
type Players []*Player

// AddPlayer add player
func (p *Players) AddPlayer(player *Player) *Player {
	PLAYERS = append(PLAYERS, player)
	return player
}

// FindPlayerWithWs find player with ws
func (p *Players) FindPlayerWithWs(conn *websocket.Conn) *Player {
	for _, player := range *p {
		if player.Conn == conn {
			return player
		}
	}
	return nil
}

// FindPlayerWithToken find player with token
func (p *Players) FindPlayerWithToken(token string) *Player {
	for _, player := range *p {
		if player.Token == token {
			return player
		}
	}
	return nil
}

// RemovePlayerWithWs delete player
func (p *Players) RemovePlayerWithWs(conn *websocket.Conn) {
	for i, player := range *p {
		if player.Conn == conn {
			*p = append((*p)[:i], (*p)[i+1:]...)
			return
		}
	}
}
