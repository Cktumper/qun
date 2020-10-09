package user

//	用户接口
//
//	Author(Wind)
type User interface {
	GetUserId() int
	GetNickName() string
}
