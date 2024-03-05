package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ret := new(Client)
	ret.Conn = conn
	return ret
}
