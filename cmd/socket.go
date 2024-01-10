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
			s.chatHandle(ws)
		case "animate":
			s.animateHandle(ws)
		default:
			HandleLog("invalid type: "+request.Type, nil)
			s.emit(ws, Response{Type: "error", Message: "invalid type: " + request.Type})
		}
		s.broadcastAll(Response{Type: "total", Message: fmt.Sprintf("total of %d rooms, total of %d players.", len(ROOMS), len(PLAYERS))})
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
func (s *Socket) broadcast(response Response) {
	for ws := range response.Room.Players {
		s.emit(ws, response)
	}
}

// broadcastAll performs an emit operation to all users.
func (s *Socket) broadcastAll(response Response) {
	for _, player := range PLAYERS {
		go func(conn *websocket.Conn) {
			s.emit(conn, response)
		}(player.Conn)
	}
}

// loginHandle If login emit is received by the client, create a new user and assign it to the room.
func (s *Socket) loginHandle(ws *websocket.Conn) {
	// check token
	if room, player := FindTokenPlayerAndRoom(request); room != nil && player != nil {
		message := fmt.Sprintf("%s: re-connected", player.Name)
		s.broadcast(Response{Type: request.Type, Message: message, Room: room, Player: player, Players: room.GetPlayers()})
		return
	}
	// create new player
	player := NewPlayer(ws)
	// return new room or exists room
	room, err := NewRoom(request)
	if err != nil || room == nil {
		PLAYERS.RemovePlayerWithWs(ws)
		s.emit(ws, Response{Type: "fatal", Message: "Failed initialized", Room: room})
		return
	}
	room.Players[ws] = player
	if len(room.Players) == 1 {
		player.IsGuessing = true
	}
	message := fmt.Sprintf("%s: connected", player.Name)
	room.AddMessage(message)
	HandleLog(message, nil)
	s.emit(ws, Response{Type: "init", Message: "Login Success", Room: room, Player: player})
	s.broadcast(Response{Type: request.Type, Message: message, Room: room, Player: player, Players: room.GetPlayers()})
}

// nameHandle If a change name emit is received by the client, change the name of the corresponding user.
func (s *Socket) nameHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoomWithWs(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s changed its name to %s", player.Name, request.Message)
		player.SetName(request.Message)
		room.AddMessage(message)
		s.broadcast(Response{Type: request.Type, Message: message, Room: room, Player: player})
	}
}

// wordleHandle word predictions operations
func (s *Socket) wordleHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoomWithWs(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("player named %s made a prediction.", player.Name)
		if room.Length == len(request.Message) {
			wordle := strings.ToUpper(request.Message)
			// If a word is used that is not in the game language, -2 points penalty. The word list is embedded in the project.
			contains, err := ContainsWord(room.Length, room.Lang, wordle)
			if !contains || err != nil {
				player.Score -= 2
				message += fmt.Sprintf("-2 points for entering a non-existent word.")
			} else {
				room.CheckWord(wordle, player)
			}
		} else {
			message = "the set word Length does not match."
		}
		room.AddMessage(message)
		s.broadcast(Response{Type: request.Type, Message: message, Room: room, Player: player})
	}
}

// chatHandle Messages intra-room correspondence
func (s *Socket) chatHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoomWithWs(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s: %s", player.Name, request.Message)
		room.AddMessage(message)
		s.broadcast(Response{Type: request.Type, Message: "new message", Room: room, Player: player})
	}
}

// animateHandle player position change on the screen
func (s *Socket) animateHandle(conn *websocket.Conn) {
	if room := ROOMS.FindRoomWithWs(conn); room != nil {
		player := room.Players[conn]
		player.Position = request.Position
		s.broadcast(Response{Type: request.Type, Message: "animate", Room: room, Player: player})
	}
}

// limitHandle game max Limit
func (s *Socket) limitHandle() error {
	if len(PLAYERS) > MaxConnection {
		return errors.New("maximum Limit reached")
	}
	return nil
}

// disconnect
func (s *Socket) disconnect(conn *websocket.Conn) {
	if room := ROOMS.FindRoomWithWs(conn); room != nil {
		player := room.Players[conn]
		message := fmt.Sprintf("%s disconnected", player.Name)
		HandleLog(message, nil)
		s.broadcast(Response{Type: "disconnect", Message: message, Room: room, Player: player})
		PLAYERS.RemovePlayerWithWs(conn)
		ROOMS.RemoveRoom(room)
		_ = conn.Close()
	}
}
