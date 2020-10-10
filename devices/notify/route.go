package notify

import (
	"encoding/json"
	"peon.top/qun/app/packet"
	"peon.top/qun/controllers/ws"
)

type (
	Handle func(sender ws.Sender, msg *packet.Packet) error

	Route struct {
		RouteItem map[uint]Handle
	}
)

func NewRoute() *Route {
	return &Route{
		RouteItem: make(map[uint]Handle),
	}
}

func (p *Route) Do(service *Service, data []byte) (err error) {
	var msg packet.Packet

	if err = json.Unmarshal(data, &msg); err != nil {
		return err
	}

	go p.Call(service, &msg)

	return nil
}

//	注册路表由函数
//
//	Author(Wind)
func (p *Route) Add(packetId uint, handle Handle) *Route {
	p.RouteItem[packetId] = handle
	return p
}

//	执行某一个路由函数
//
//	Author(Wind)
func (p *Route) Call(service *Service, msg *packet.Packet) {
	if handle, ok := p.RouteItem[msg.PacketId]; ok {
		_ = handle(service, msg)
	}
}
