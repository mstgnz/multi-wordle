package main

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

var request Request

type Socket struct{}

func NewSocket() *Socket {
	return &Socket{}
}

// Handler handle incoming requests from the client.
func (s *Socket) Handler(ws *websocket.Conn) {

	defer func(ws *websocket.Conn) {
		s.disconnect(ws)
	}(ws)

	if err := s.limitHandle(ws); err != nil {
		s.emit(ws, Response{Type: "error", Message: err.Error()})
		return
	}

	for {
		if err := websocket.JSON.Receive(ws, &request); err != nil {
			if err == io.EOF {
				break
			}
			HandleLog("read error", err)
			s.emit(ws, Response{Type: "error", Message: err.Error()})
			continue
		}
		switch request.Type {
		case "login":
			s.loginHandle(ws)
		case "change-name":
			s.changeNameHandle(ws)
		case "wordle":
			s.wordleHandle(ws)
		case "chat":
			s.messageHandle(ws)
		default:
			HandleLog("invalid type: "+request.Type, nil)
			s.emit(ws, Response{Type: "error", Message: "invalid type: " + request.Type})
		}
	}
}

// emit Performs emit operation to the user taking action
func (s *Socket) emit(ws *websocket.Conn, response Response) {
	err := websocket.JSON.Send(ws, response)
	if err != nil {
		HandleLog("emit error", err)
	}
}

// broadcast The user receiving the action performs an emit operation to the user in the room to which it belongs.
func (s *Socket) broadcast(conn *websocket.Conn, response Response) {
	for ws := range ROOMS.FindRoom(conn).Players {
		s.emit(ws, response)
	}
}

// loginHandle If login emit is received by the client, create a new user and assign it to the room.
func (s *Socket) loginHandle(ws *websocket.Conn) {
	player := NewPlayer(ws)
	room := NewRoom("en", 5)
	room.Players[ws] = player
	HandleLog(player.Name+" connected", nil)
	s.broadcast(ws, Response{Type: request.Type, Message: "new player connected", Room: room})
}

// changeNameHandle If a change name emit is received by the client, change the name of the corresponding user.
func (s *Socket) changeNameHandle(conn *websocket.Conn) {
	if player := PLAYERS.FindPlayer(conn); player != nil {
		message := fmt.Sprintf("%s changed its name to %s", player.Name, request.Message)
		player.SetName(request.Message)
		s.broadcast(conn, Response{Type: request.Type, Message: message})
	}
}

// wordleHandle word predictions operations
func (s *Socket) wordleHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		room.Players[conn].AddWordToGuess(request.Message)
		room.Wordle.CheckWord(request.Message)
		s.broadcast(conn, Response{Type: request.Type, Message: "new wordle", Room: room})
	}
}

// messageHandle Messages intra-room correspondence
func (s *Socket) messageHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s: %s", player.Name, request.Message)
		room.AddMessage(message)
		s.broadcast(conn, Response{Type: request.Type, Message: "new message", Room: room})
	}
}

// limitHandle game max limit
func (s *Socket) limitHandle(_ *websocket.Conn) error {
	if len(PLAYERS) > MaxConnection {
		return errors.New("maximum limit reached")
	}
	return nil
}

// disconnect
func (s *Socket) disconnect(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		msg := fmt.Sprintf("%s disconnected", player.Name)
		HandleLog(msg, nil)
		PLAYERS.DelPlayer(conn)
		s.broadcast(conn, Response{Type: "disconnect", Message: msg, Room: room})
		_ = conn.Close()
	}
}
