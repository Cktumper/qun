package kernel

//	网络连接抽象接口
//
//	Author(Wind)
type Connection interface {
	//	发送消息
	Send(msg []byte) error

	//	接收连接信息
	//	返回接收的信息 该信息则直接会被发送给客户端
	//	返回是否回信，bool 为 false 则不回信
	//	如果出错则会直接关闭消息循环
	Receiver() (Packet, bool, error)

	//	关闭连接
	Close()
}
