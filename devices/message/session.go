package message

import "github.com/gorilla/websocket"

type (
	Connection interface {
		GetConnection() *websocket.Conn
	}

	Session interface {
		Connection
		GetSessionId() string
	}
)
