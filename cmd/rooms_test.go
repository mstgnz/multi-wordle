package main

import (
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
		{
			name: "İki kişilik oda ekleme testi",
			r:    Rooms{},
			args: args{
				room: &Room{
					Name:    "Test Room",
					Lang:    "tr",
					Limit:   2,
					Length:  5,
					Start:   false,
					Players: make(map[*websocket.Conn]*Player),
				},
			},
			want: &Room{
				Name:    "Test Room",
				Lang:    "tr",
				Limit:   2,
				Length:  5,
				Start:   false,
				Players: make(map[*websocket.Conn]*Player),
			},
		},
		{
			name: "Tek kişilik oda ekleme testi",
			r:    Rooms{},
			args: args{
				room: &Room{
					Name:    "Single Room",
					Lang:    "tr",
					Limit:   1,
					Length:  5,
					Start:   false,
					Players: make(map[*websocket.Conn]*Player),
				},
			},
			want: &Room{
				Name:    "Single Room",
				Lang:    "tr",
				Limit:   1,
				Length:  5,
				Start:   false,
				Players: make(map[*websocket.Conn]*Player),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.AddRoom(tt.args.room)
			if got.Name != tt.want.Name || got.Lang != tt.want.Lang || got.Limit != tt.want.Limit {
				t.Errorf("AddRoom() = %v, want %v", got, tt.want)
			}
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
		{
			name: "Websocket bağlantısı ile oda bulma testi",
			r: Rooms{
				&Room{
					Name: "Test Room",
					Players: map[*websocket.Conn]*Player{
						{}: {},
					},
				},
			},
			args: args{
				conn: &websocket.Conn{},
			},
			want: &Room{
				Name: "Test Room",
				Players: map[*websocket.Conn]*Player{
					{}: {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.FindRoomWithWs(tt.args.conn)
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("FindRoomWithWs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRooms_RemoveRoom(t *testing.T) {
	type args struct {
		room *Room
	}
	tests := []struct {
		name string
		r    Rooms
		args args
		want bool
	}{
		{
			name: "Oda silme testi",
			r: Rooms{
				&Room{
					Name:    "Test Room",
					Players: make(map[*websocket.Conn]*Player),
				},
			},
			args: args{
				room: &Room{
					Name:    "Test Room",
					Players: make(map[*websocket.Conn]*Player),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialLen := len(tt.r)
			tt.r.RemoveRoom(tt.args.room)
			if len(tt.r) != initialLen-1 {
				t.Errorf("RemoveRoom() failed, room was not removed")
			}
		})
	}
}
