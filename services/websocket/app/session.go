package app

import (
	"log"
	"peon.top/qun/chat"
)

type Session struct {
	user       *User
	connection *Connection
}

//	创建一个 Session 对象
//
//	Author(Wind)
func NewSession(user *User, connection *Connection) *Session {
	return &Session{
		user:       user,
		connection: connection,
	}
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
func (p *Session) GetConnection() chat.Connection {
	return p.connection
}

//	用户离开事件
//
//	Author(Wind)
func (p *Session) OnLeafed(room *chat.Room) {
	log.Printf("%s 离开了 %s", p.GetUserInformation().GetNickname(), room.Name)
}
