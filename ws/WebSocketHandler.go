package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	bufferSize = 2048
)

type WebSocketHandler struct {
	Server *Server
}

func NewWebSocketHandler() *WebSocketHandler {
	ret := new(WebSocketHandler)
	ret.Server = NewServer()
	go ret.Server.Run()
	return ret
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *WebSocketHandler) get(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{ReadBufferSize: bufferSize, WriteBufferSize: bufferSize}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "upgrade error", http.StatusInternalServerError)
	}
	client := NewClient(conn)
	h.Server.Clients[client] = true
}
