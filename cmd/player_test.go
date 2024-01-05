package main

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/websocket"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
		want *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := NewPlayer(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_SetIsGuessing(t *testing.T) {
	type fields struct {
		Name       string
		Score      int
		IsGuessing bool
		Color      string
		Position   Position
		Conn       *websocket.Conn
	}
	type args struct {
		isGuessing bool
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
			p := &Player{
				Name:       tt.fields.Name,
				Score:      tt.fields.Score,
				IsGuessing: tt.fields.IsGuessing,
				Color:      tt.fields.Color,
				Position:   tt.fields.Position,
				Conn:       tt.fields.Conn,
			}
			p.SetIsGuessing(tt.args.isGuessing)
		})
	}
}

func TestPlayer_SetName(t *testing.T) {
	type fields struct {
		Name       string
		Score      int
		IsGuessing bool
		Color      string
		Position   Position
		Conn       *websocket.Conn
	}
	type args struct {
		name string
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
			p := &Player{
				Name:       tt.fields.Name,
				Score:      tt.fields.Score,
				IsGuessing: tt.fields.IsGuessing,
				Color:      tt.fields.Color,
				Position:   tt.fields.Position,
				Conn:       tt.fields.Conn,
			}
			p.SetName(tt.args.name)
		})
	}
}

func TestPlayer_SetScore(t *testing.T) {
	type fields struct {
		Name       string
		Score      int
		IsGuessing bool
		Color      string
		Position   Position
		Conn       *websocket.Conn
	}
	type args struct {
		score int
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
			p := &Player{
				Name:       tt.fields.Name,
				Score:      tt.fields.Score,
				IsGuessing: tt.fields.IsGuessing,
				Color:      tt.fields.Color,
				Position:   tt.fields.Position,
				Conn:       tt.fields.Conn,
			}
			p.SetScore(tt.args.score)
		})
	}
}
