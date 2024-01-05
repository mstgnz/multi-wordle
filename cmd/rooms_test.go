package main

import (
	"reflect"
	"testing"

	"golang.org/x/net/websocket"
)

func TestRooms_AddRoom(t *testing.T) {
	type args struct {
		room *Room
	}
	tests := []struct {
		name string
		r    Rooms
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddRoom(tt.args.room); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRooms_DelRoom(t *testing.T) {
	tests := []struct {
		name string
		r    Rooms
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.DelRoom()
		})
	}
}

func TestRooms_FindRoom(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		name string
		r    Rooms
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FindRoom(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
