package main

import (
	"reflect"
	"strconv"
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
		args    args
		want    *Room
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := NewRoom(tt.args.lang, tt.args.length, tt.args.trial)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoom() got = %v, want %v", got, tt.want)
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
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := &Room{
				ID:         tt.fields.ID,
				Length:     tt.fields.Length,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
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
		fields fields
		want   []*Player
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := &Room{
				ID:         tt.fields.ID,
				Length:     tt.fields.Length,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			if got := r.GetPlayers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}
