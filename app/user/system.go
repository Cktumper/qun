package user

type System struct{}

//	构建系统用户
//
//	Author(Wind)
func NewSystem() *System {
	return &System{}
}

//	获取用户ID
//
//	Author(Wind)
func (p *System) GetUserId() int {
	return 0
}

//	系统用户
//
//	Author(Wind)
func (p *System) GetNickName() string {
	return "系统"
}

//	是否被阻挡
//
//	Author(Wind)
func (p *System) Blocked() bool {
	return false
}
