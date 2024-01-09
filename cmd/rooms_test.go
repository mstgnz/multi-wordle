package main

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/websocket"
)

func TestRooms_AddRoom(t *testing.T) {
	type args struct {
		room *Room
	}
	tests := []struct {
		r    Rooms
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := tt.r.AddRoom(tt.args.room); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRooms_DelRoom(t *testing.T) {
	tests := []struct {
		r Rooms
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.r.RemoveRoom()
		})
	}
}

func TestRooms_FindRoom(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		r    Rooms
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := tt.r.FindRoomWithWs(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRoomWithWs() = %v, want %v", got, tt.want)
			}
		})
	}
}
