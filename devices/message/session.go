package message

import "github.com/gorilla/websocket"

type Connection interface {
	GetConnection() *websocket.Conn
}
