package kernel

type User interface {
	//	获取用户ID
	GetUserId() int

	//	获取用户昵称
	GetNickname() string
}
