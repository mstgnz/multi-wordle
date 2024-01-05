package main

import (
	"reflect"
	"sync"
	"testing"

	"golang.org/x/net/websocket"
)

func TestNewRoom(t *testing.T) {
	type args struct {
		lang   string
		length int
		trial  int
	}
	tests := []struct {
		name string
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoom(tt.args.lang, tt.args.length, tt.args.trial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRoom1(t *testing.T) {
	type args struct {
		lang   string
		length int
		trial  int
	}
	tests := []struct {
		name string
		args args
		want *Room
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoom(tt.args.lang, tt.args.length, tt.args.trial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_AddMessage(t *testing.T) {
	type fields struct {
		ID         string
		Length     int
		Trial      int
		Messages   []string
		Wordle     Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
		Mutex      sync.Mutex
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				ID:         tt.fields.ID,
				Length:     tt.fields.Length,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
				Mutex:      tt.fields.Mutex,
			}
			r.AddMessage(tt.args.message)
		})
	}
}

func TestRoom_GetPlayers(t *testing.T) {
	type fields struct {
		ID         string
		Length     int
		Trial      int
		Messages   []string
		Wordle     Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
		Mutex      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Player
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				ID:         tt.fields.ID,
				Length:     tt.fields.Length,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
				Mutex:      tt.fields.Mutex,
			}
			if got := r.GetPlayers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}
