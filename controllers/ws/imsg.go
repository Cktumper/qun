package ws

type IMsg interface {
	GetFrom() interface{}
	GetTo() interface{}
	GetPacket() interface{}
}
