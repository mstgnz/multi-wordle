package main

import (
	"golang.org/x/net/websocket"
)

// Rooms We keep the room list as a pointer to facilitate room operations
type Rooms []*Room

// AddRoom add room
func (r *Rooms) AddRoom(room *Room) *Room {
	ROOMS = append(ROOMS, room)
	return room
}

// FindRoomWithWs find the room by connection
func (r *Rooms) FindRoomWithWs(conn *websocket.Conn) *Room {
	for _, room := range *r {
		if _, ok := room.Players[conn]; ok {
			return room
		}
	}
	return nil
}

// RemoveRoom delete the room if there are no players in the room
func (r *Rooms) RemoveRoom(room *Room) {
	for i, rm := range *r {
		if rm == room && len(rm.Players) == 0 {
			*r = append((*r)[:i], (*r)[i+1:]...)
			return
		}
	}
}
