package main

import (
	"reflect"
	"testing"

	"golang.org/x/net/websocket"
)

func TestPlayers_AddPlayer(t *testing.T) {
	type args struct {
		player *Player
	}
	tests := []struct {
		name string
		p    Players
		args args
		want *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.AddPlayer(tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayers_DelPlayer(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		name string
		p    Players
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DelPlayer(tt.args.conn)
		})
	}
}

func TestPlayers_FindPlayer(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		name string
		p    Players
		args args
		want *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.FindPlayer(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
