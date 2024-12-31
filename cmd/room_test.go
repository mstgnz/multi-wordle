package main

import (
	"strconv"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

func TestNewRoom(t *testing.T) {
	type args struct {
		request Request
	}
	tests := []struct {
		name    string
		args    args
		want    *Room
		wantErr bool
	}{
		{
			name: "İki kişilik oda oluşturma testi",
			args: args{
				request: Request{
					Lang:   "tr",
					Limit:  2,
					Length: 5,
				},
			},
			want: &Room{
				Lang:     "tr",
				Limit:    2,
				Length:   5,
				Start:    false,
				Trial:    0,
				Messages: []string{},
				Players:  make(map[*websocket.Conn]*Player),
			},
			wantErr: false,
		},
		{
			name: "Tek kişilik oda oluşturma testi",
			args: args{
				request: Request{
					Lang:   "tr",
					Limit:  1,
					Length: 5,
				},
			},
			want: &Room{
				Lang:     "tr",
				Limit:    1,
				Length:   5,
				Start:    false,
				Trial:    0,
				Messages: []string{},
				Players:  make(map[*websocket.Conn]*Player),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoom(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil {
				if got.Lang != tt.want.Lang || got.Limit != tt.want.Limit || got.Length != tt.want.Length {
					t.Errorf("NewRoom() = %v, want %v", got, tt.want)
				}
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
		name   string
		fields fields
		args   args
	}{
		{
			name: "Doğru kelime testi - 10 puan",
			fields: fields{
				Wordle: &Wordle{Word: "KALEM"},
				Start:  true,
			},
			args: args{
				word:   "KALEM",
				player: &Player{Score: 0},
			},
		},
		{
			name: "Doğru harf doğru konum testi - 5 puan",
			fields: fields{
				Wordle: &Wordle{Word: "KALEM"},
				Start:  true,
			},
			args: args{
				word:   "KADIN",
				player: &Player{Score: 0},
			},
		},
		{
			name: "Doğru harf yanlış konum testi - 3 puan",
			fields: fields{
				Wordle: &Wordle{Word: "KALEM"},
				Start:  true,
			},
			args: args{
				word:   "EKRAN",
				player: &Player{Score: 0},
			},
		},
		{
			name: "Olmayan harf ikinci kullanım testi - -1 puan",
			fields: fields{
				Wordle: &Wordle{Word: "KALEM"},
				Start:  true,
			},
			args: args{
				word:   "ZZZZT",
				player: &Player{Score: 0, UsedLetters: map[string]bool{"Z": true}},
			},
		},
		{
			name: "Geçersiz kelime testi - -2 puan",
			fields: fields{
				Wordle: &Wordle{Word: "KALEM"},
				Start:  true,
				Lang:   "tr",
			},
			args: args{
				word:   "XXXXX",
				player: &Player{Score: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

		want []*Player
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
			got := r.GetPlayers()
			if len(got) != len(tt.want) {
				t.Errorf("GetPlayers() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i, player := range got {
				if player.Name != tt.want[i].Name || player.Score != tt.want[i].Score {
					t.Errorf("GetPlayers()[%d] = %v, want %v", i, player, tt.want[i])
				}
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
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Süre aşımı testi - -5 puan",
			fields: fields{
				Start:      true,
				PlayerTurn: &websocket.Conn{},
				Players: map[*websocket.Conn]*Player{
					{}: {
						LastActivity: time.Now().Add(-31 * time.Second),
						Score:        0,
					},
				},
			},
			args: args{
				conn: &websocket.Conn{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			if got != nil && tt.want != nil {
				if got.Lang != tt.want.Lang || got.Limit != tt.want.Limit || got.Length != tt.want.Length || got.Start != tt.want.Start {
					t.Errorf("NextMatch() = %v, want %v", got, tt.want)
				}
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
			if got != nil && tt.want != nil {
				if got.Lang != tt.want.Lang || got.Limit != tt.want.Limit || got.Length != tt.want.Length || got.Start != tt.want.Start {
					t.Errorf("ResetMatch() = %v, want %v", got, tt.want)
				}
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
		name   string
		fields fields
		want   *Player
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			got := r.FindGuessing()
			if got != nil && tt.want != nil {
				if got.Name != tt.want.Name || got.Score != tt.want.Score {
					t.Errorf("FindGuessing() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
