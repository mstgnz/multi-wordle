package main

import (
	"errors"
	"io"

	"golang.org/x/net/websocket"
)

var request Request

type Socket struct {
	connections map[*websocket.Conn]bool
	Room
}

func NewSocket() *Socket {
	return &Socket{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Socket) Handler(ws *websocket.Conn) {

	defer func(ws *websocket.Conn) {
		s.disconnect(ws)
	}(ws)

	if err := s.limitHandle(ws); err != nil {
		s.emit(ws, Response{Type: "error", Message: err.Error()})
		delete(s.connections, ws)
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
		s.connections[ws] = true
		switch request.Type {
		case "new":
			s.newHandle(ws)
		case "message":
			s.messageHandle(ws)
		case "animate":
			s.animateHandle(ws)
		case "name":
			s.nameHandle(ws)
		default:
			HandleLog("invalid type: "+request.Type, nil)
			s.emit(ws, Response{Type: "error", Message: "invalid type: " + request.Type})
		}
	}
}

func (s *Socket) emit(ws *websocket.Conn, response Response) {
	err := websocket.JSON.Send(ws, response)
	if err != nil {
		HandleLog("emit error", err)
	}
}

func (s *Socket) broadcast(response Response) {
	for ws := range s.connections {
		if s.connections[ws] {
			// TODO to many connection handle
			go func(ws *websocket.Conn) {
				s.emit(ws, response)
			}(ws)
		}
	}
}

func (s *Socket) newHandle(ws *websocket.Conn) {
	/*player := players.AddPlayer(types.Player{
		Color:    RandomColor(),
		Name:     request.Player.Name,
		Position: types.Position{X: 0, Y: 0},
	})
	HandleLog(request.Player.Name+" connected", nil)
	s.emit(ws, Response{Type: "init", Message: "login successfully", Player: player, Players: players, Messages: messages})
	s.broadcast(Response{Type: request.Type, Message: "new player connected", Player: player, Messages: messages})*/
}

func (s *Socket) messageHandle(_ *websocket.Conn) {
	/*messages = append(messages, types.Message{Name: request.Player.Name, Message: request.Message})
	if player := players.FindPlayer(request.Player.Name); player != nil {
		s.broadcast(Response{Type: request.Type, Message: request.Message, Player: *player})
	}*/
}

func (s *Socket) animateHandle(_ *websocket.Conn) {
	/*if player := players.FindPlayer(request.Player.Name); player != nil {
		player.Position = request.Player.Position
		s.broadcast(Response{Type: request.Type, Message: "animate", Player: *player})
	}*/
}

func (s *Socket) nameHandle(_ *websocket.Conn) {
	/*if player := players.FindPlayer(request.Player.Name); player != nil {
		s.broadcast(Response{Type: request.Type, Message: request.Message, Player: *player})
		player.Name = request.Message
	}*/
}

func (s *Socket) limitHandle(_ *websocket.Conn) error {
	if len(s.connections) > 50 {
		return errors.New("maximum limit reached")
	}
	return nil
}

func (s *Socket) disconnect(ws *websocket.Conn) {
	/*msg := fmt.Sprintf("%s disconnected", request.Player.Name)
	HandleLog(msg, nil)
	s.connections[ws] = false
	players.DelPlayer(request.Player.Name)
	delete(s.connections, ws)
	s.broadcast(Response{
		Type:    "disconnect",
		Message: msg,
		Player:  request.Player,
	})*/
	_ = ws.Close()
}
