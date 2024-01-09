package main

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/websocket"
)

func TestPlayers_AddPlayer(t *testing.T) {
	type args struct {
		player *Player
	}
	tests := []struct {
		p    Players
		args args
		want *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
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
		p    Players
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.p.RemovePlayerWithWs(tt.args.conn)
		})
	}
}

func TestPlayers_FindPlayer(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		p    Players
		args args
		want *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := tt.p.FindPlayerWithWs(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPlayerWithWs() = %v, want %v", got, tt.want)
			}
		})
	}
}
