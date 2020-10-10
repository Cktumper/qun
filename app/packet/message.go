package packet

type Message struct {
	From   interface{}
	To     interface{}
	Packet interface{}
}

func (p *Message) GetFrom() interface{} {
	return p.From
}

func (p *Message) GetTo() interface{} {
	return p.To
}

func (p *Message) GetPacket() interface{} {
	return p.Packet
}
