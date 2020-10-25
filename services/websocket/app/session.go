package app

import "peon.top/qun/chat"

type Session struct {
	user       *User
	connection *Connection
}

//	获取用户信息
//
//	Author(Wind)
func (p *Session) GetUserInformation() chat.User {
	return p.user
}

//	获取用户连接
//
//	Author(Wind)
func (p *Session) GetCollection() chat.Connection {
	return p.connection
}
