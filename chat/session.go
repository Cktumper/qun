package chat

type Session interface {
	//	获取网络连接
	//	用于接收发送信息
	GetConnection() Connection

	//	获取用户信息
	GetUserInformation() User

	//	Session 的关闭，表示的离开
	//	离开事件
	OnLeafed()
}
