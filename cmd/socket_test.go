package main

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/websocket"
)

func TestNewSocket(t *testing.T) {
	tests := []struct {
		want *Socket
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := NewSocket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocket_Handler(t *testing.T) {
	type args struct {
		ws *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.Handler(tt.args.ws)
		})
	}
}

func TestSocket_animateHandle(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.animateHandle(tt.args.conn)
		})
	}
}

func TestSocket_broadcast(t *testing.T) {
	type args struct {
		conn     *websocket.Conn
		response Response
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.broadcast(tt.args.conn, tt.args.response)
		})
	}
}

func TestSocket_disconnect(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.disconnect(tt.args.conn)
		})
	}
}

func TestSocket_emit(t *testing.T) {
	type args struct {
		ws       *websocket.Conn
		response Response
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.emit(tt.args.ws, tt.args.response)
		})
	}
}

func TestSocket_limitHandle(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			if err := s.limitHandle(); (err != nil) != tt.wantErr {
				t.Errorf("limitHandle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSocket_loginHandle(t *testing.T) {
	type args struct {
		ws *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.loginHandle(tt.args.ws)
		})
	}
}

func TestSocket_messageHandle(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.messageHandle(tt.args.conn)
		})
	}
}

func TestSocket_nameHandle(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.nameHandle(tt.args.conn)
		})
	}
}

func TestSocket_wordleHandle(t *testing.T) {
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := &Socket{}
			s.wordleHandle(tt.args.conn)
		})
	}
}
