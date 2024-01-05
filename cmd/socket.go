package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

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

	if err := s.limitHandle(); err != nil {
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
		case "name":
			s.nameHandle(ws)
		case "wordle":
			s.wordleHandle(ws)
		case "chat":
			s.messageHandle(ws)
		case "animate":
			s.animateHandle(ws)
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
	room, err := NewRoom("en", 5, 5)
	if err != nil || room == nil {
		PLAYERS.DelPlayer(ws)
		s.emit(ws, Response{Type: request.Type, Message: "Failed initialized", Room: room})
		return
	}
	if len(room.Players) == 0 {
		player.IsGuessing = true
	}
	room.Players[ws] = player
	message := fmt.Sprintf("%s: connected", player.Name)
	HandleLog(message, nil)
	s.broadcast(ws, Response{Type: request.Type, Message: message, Room: room, Player: player, Players: room.GetPlayers()})
}

// nameHandle If a change name emit is received by the client, change the name of the corresponding user.
func (s *Socket) nameHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s changed its name to %s", player.Name, request.Message)
		player.SetName(request.Message)
		s.broadcast(conn, Response{Type: request.Type, Message: message, Room: room, Player: player, Players: room.GetPlayers()})
	}
}

// wordleHandle word predictions operations
func (s *Socket) wordleHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		if room.Length == len(request.Message) {
			room.Wordle.CheckWord(strings.ToUpper(request.Message))
			s.broadcast(conn, Response{Type: request.Type, Message: "new wordle", Room: room, Player: player, Players: room.GetPlayers()})
		} else {
			s.emit(conn, Response{Type: "error", Message: "word count not matched", Room: room, Player: player, Players: room.GetPlayers()})
		}
	}
}

// messageHandle Messages intra-room correspondence
func (s *Socket) messageHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s: %s", player.Name, request.Message)
		room.AddMessage(message)
		s.broadcast(conn, Response{Type: request.Type, Message: "new message", Room: room, Player: player, Players: room.GetPlayers()})
	}
}

// animateHandle player position change on the screen
func (s *Socket) animateHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoom(conn); room != nil {
		player := room.Players[conn]
		player.Position = request.Position
		s.broadcast(conn, Response{Type: request.Type, Message: "animate", Room: room, Player: player, Players: room.GetPlayers()})
	}
}

// limitHandle game max limit
func (s *Socket) limitHandle() error {
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
		s.broadcast(conn, Response{Type: "disconnect", Message: msg, Room: room, Player: player, Players: room.GetPlayers()})
		_ = conn.Close()
	}
}
