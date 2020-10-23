package kernel

type Session interface {
	//	获取网络连接
	//	用于接收发送信息
	GetConnection() Connection

	//	获取用户信息
	GetUserInformation() User
}
