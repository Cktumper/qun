package app

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"peon.top/qun/chat"
)

type Connection struct {
	conn *websocket.Conn
}

//	构建新的连接器
//
//	Author(Wind)
func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{conn: conn}
}

//	发送消息
//
//	Author(Wind)
func (p *Connection) Send(message chat.Packet) error {
	return p.conn.WriteMessage(1, message.Marshal())
}

//	接收消息
//
//	Author(Wind)
func (p *Connection) Receiver() (chat.Packet, bool, error) {
	_, message, err := p.conn.ReadMessage()

	//	告诉核心，出错了
	if err != nil {
		return nil, false, err
	}

	//	处理消息包
	//	构建包结构
	packet := Packet{}

	//	解析包
	if err := json.Unmarshal(message, &packet); err != nil {
		return nil, false, err
	}

	//	一定会有回信
	return &packet.PacketText, true, nil
}

//	关闭连接
//
//	Author(Wind)
func (p *Connection) Close() {
	_ = p.conn.Close()
}
