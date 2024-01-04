package main

import (
	"golang.org/x/net/websocket"
)

// Rooms We keep the room list as a pointer to facilitate room operations
type Rooms []*Room

// AddRoom add room
func (r *Rooms) AddRoom(room *Room) *Room {
	*r = append(*r, room)
	return room
}

// FindRoom find the room by connection
func (r *Rooms) FindRoom(conn *websocket.Conn) *Room {
	for _, room := range *r {
		if _, ok := room.Players[conn]; ok {
			return room
		}
	}
	return nil
}

// DelRoom delete the room if there are no players in the room
func (r *Rooms) DelRoom() {
	for i, room := range *r {
		if len(room.Players) == 0 {
			*r = append((*r)[:i], (*r)[i+1:]...)
			return
		}
	}
}
