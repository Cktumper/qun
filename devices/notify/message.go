package notify

type Message interface {
	GetFrom() interface{}
	GetTo() interface{}
	GetPacket() interface{}
}
