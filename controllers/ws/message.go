package ws

import (
	"fmt"
	"peon.top/qun/app/packet"
	"peon.top/qun/devices/session"
	"sync"
)

type MessagePacketController struct{}

var (
	messagePacketInstance     *MessagePacketController
	messagePacketOnceInstance sync.Once
)

func NewMessagePacketController() *MessagePacketController {
	messagePacketOnceInstance.Do(func() {
		messagePacketInstance = &MessagePacketController{}
	})
	return messagePacketInstance
}

func (p *MessagePacketController) Chat(service Sender, msg *packet.Packet) error {
	fmt.Println(msg.Chat.Message)

	self, _ := session.GetService().Find(msg.SessionId)

	service.Send(&packet.Message{
		From: self,
		To:   session.GetService().GetAll(nil),
		Packet: &packet.Chat{
			Base:    packet.Base{PacketId: 2},
			Message: msg.Chat.Message,
		},
	})

	return nil
}
