package app

//	用户结构
//
//	Author(Wind)
type User struct {
	UserId   int
	Nickname string
}

//	初始化用户对象
//
//	Author(Wind)
func NewUser(nickname string) *User {
	return &User{
		UserId:   1,
		Nickname: nickname,
	}
}

//	获取用户ID
//
//	Author(Wind)
func (p *User) GetUserId() int {
	return p.UserId
}

//	获取用户昵称
//
//	Author(Wind)
func (p *User) GetNickname() string {
	if len(p.Nickname) <= 0 {
		return "匿名"
	}

	return p.Nickname
}
