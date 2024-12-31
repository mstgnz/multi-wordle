package main

import (
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
		{
			name: "Yeni oyuncu ekleme testi",
			p:    Players{},
			args: args{
				player: &Player{
					Name:       "TestPlayer",
					Score:      0,
					IsGuessing: false,
				},
			},
			want: &Player{
				Name:       "TestPlayer",
				Score:      0,
				IsGuessing: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.AddPlayer(tt.args.player)
			if got != nil && tt.want != nil {
				if got.Name != tt.want.Name || got.Score != tt.want.Score || got.IsGuessing != tt.want.IsGuessing {
					t.Errorf("AddPlayer() = %v, want %v", got, tt.want)
				}
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
		{
			name: "Oyuncu silme testi",
			p: Players{
				{
					Name:       "TestPlayer",
					Score:      0,
					IsGuessing: false,
					Conn:       &websocket.Conn{},
				},
			},
			args: args{
				conn: &websocket.Conn{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialLen := len(tt.p)
			tt.p.RemovePlayerWithWs(tt.args.conn)
			if len(tt.p) != initialLen-1 {
				t.Errorf("RemovePlayerWithWs() failed, player was not removed")
			}
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
		{
			name: "Oyuncu bulma testi",
			p: Players{
				{
					Name:       "TestPlayer",
					Score:      10,
					IsGuessing: true,
					Conn:       &websocket.Conn{},
				},
			},
			args: args{
				conn: &websocket.Conn{},
			},
			want: &Player{
				Name:       "TestPlayer",
				Score:      10,
				IsGuessing: true,
				Conn:       &websocket.Conn{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.FindPlayerWithWs(tt.args.conn)
			if got != nil && tt.want != nil {
				if got.Name != tt.want.Name || got.Score != tt.want.Score || got.IsGuessing != tt.want.IsGuessing {
					t.Errorf("FindPlayerWithWs() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
