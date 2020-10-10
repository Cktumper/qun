package ws

type Sender interface {
	Send(msg IMsg)
}
