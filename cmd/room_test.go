package main

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/websocket"
)

func TestNewRoom(t *testing.T) {
	type args struct {
		request Request
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
			got, err := NewRoom(tt.args.request)
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
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
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
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			r.AddMessage(tt.args.message)
		})
	}
}

func TestRoom_CheckWord(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	type args struct {
		word   string
		player *Player
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
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			r.CheckWord(tt.args.word, tt.args.player)
		})
	}
}

func TestRoom_GetPlayers(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
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
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			if got := r.GetPlayers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_NextGuessing(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	type args struct {
		conn *websocket.Conn
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
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			r.NextGuessing(tt.args.conn)
		})
	}
}

func TestRoom_NextMatch(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	tests := []struct {
		fields  fields
		want    *Room
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := &Room{
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			got, err := r.NextMatch()
			if (err != nil) != tt.wantErr {
				t.Errorf("NextMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_RemovePlayer(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	type args struct {
		conn *websocket.Conn
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
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			r.RemovePlayer(tt.args.conn)
		})
	}
}

func TestRoom_ResetMatch(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	tests := []struct {
		fields  fields
		want    *Room
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := &Room{
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			got, err := r.ResetMatch()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResetMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_FindGuessing(t *testing.T) {
	type fields struct {
		Name       string
		Lang       string
		Limit      int
		Length     int
		Start      bool
		Trial      int
		Messages   []string
		Wordle     *Wordle
		Matches    []*Wordle
		Players    map[*websocket.Conn]*Player
		PlayerTurn *websocket.Conn
	}
	tests := []struct {
		fields fields
		want   *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := &Room{
				Name:       tt.fields.Name,
				Lang:       tt.fields.Lang,
				Limit:      tt.fields.Limit,
				Length:     tt.fields.Length,
				Start:      tt.fields.Start,
				Trial:      tt.fields.Trial,
				Messages:   tt.fields.Messages,
				Wordle:     tt.fields.Wordle,
				Matches:    tt.fields.Matches,
				Players:    tt.fields.Players,
				PlayerTurn: tt.fields.PlayerTurn,
			}
			if got := r.FindGuessing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindGuessing() = %v, want %v", got, tt.want)
			}
		})
	}
}
