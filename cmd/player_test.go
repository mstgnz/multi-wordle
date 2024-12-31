package main

import (
	"strconv"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		name string
		args args
		want *Player
	}{
		{
			name: "Yeni oyuncu oluşturma testi",
			args: args{
				conn: &websocket.Conn{},
			},
			want: &Player{
				Name:         "",
				Score:        0,
				IsGuessing:   false,
				Color:        "",
				Position:     Position{},
				Conn:         &websocket.Conn{},
				LastActivity: time.Now(),
				UsedLetters:  make(map[string]bool),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlayer(tt.args.conn)
			if got.Name != tt.want.Name || got.Score != tt.want.Score || got.IsGuessing != tt.want.IsGuessing {
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
		Name         string
		Score        int
		IsGuessing   bool
		Color        string
		Position     Position
		Conn         *websocket.Conn
		LastActivity time.Time
		UsedLetters  map[string]bool
	}
	type args struct {
		score int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Doğru kelime puanı - 10 puan",
			fields: fields{
				Score: 0,
			},
			args: args{
				score: 10,
			},
			want: 10,
		},
		{
			name: "Doğru harf doğru konum puanı - 5 puan",
			fields: fields{
				Score: 0,
			},
			args: args{
				score: 5,
			},
			want: 5,
		},
		{
			name: "Doğru harf yanlış konum puanı - 3 puan",
			fields: fields{
				Score: 0,
			},
			args: args{
				score: 3,
			},
			want: 3,
		},
		{
			name: "Olmayan harf ikinci kullanım cezası - -1 puan",
			fields: fields{
				Score: 5,
			},
			args: args{
				score: -1,
			},
			want: 4,
		},
		{
			name: "Geçersiz kelime cezası - -2 puan",
			fields: fields{
				Score: 5,
			},
			args: args{
				score: -2,
			},
			want: 3,
		},
		{
			name: "Süre aşımı cezası - -5 puan",
			fields: fields{
				Score: 5,
			},
			args: args{
				score: -5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name:         tt.fields.Name,
				Score:        tt.fields.Score,
				IsGuessing:   tt.fields.IsGuessing,
				Color:        tt.fields.Color,
				Position:     tt.fields.Position,
				Conn:         tt.fields.Conn,
				LastActivity: tt.fields.LastActivity,
				UsedLetters:  tt.fields.UsedLetters,
			}
			p.SetScore(tt.args.score)
			if p.Score != tt.want {
				t.Errorf("SetScore() = %v, want %v", p.Score, tt.want)
			}
		})
	}
}
