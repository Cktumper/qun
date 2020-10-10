package message

import "peon.top/qun/app/packet"

//	定义一条消息
//
//	Author(Wind)
type Message struct {
	//	发送与接收用户
	From Session
	To   []Session

	//	消息包
	packet packet.Packet
}

//	构建新的消息
//
//	Author(Wind)
func NewMessage(from Session, to []Session, packet packet.Packet) *Message {
	return &Message{
		From:   from,
		To:     to,
		packet: packet,
	}
}
