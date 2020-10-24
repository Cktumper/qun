package kernel

import "sync"

type Service struct {
	container *RoomContainer
}

var (
	serviceInstance     *Service
	serviceOnceInstance sync.Once
)

//	初始化消息核心对象
//
//	Author(Wind)
func NewService() *Service {
	serviceOnceInstance.Do(func() {
		serviceInstance = &Service{}
	})

	return serviceInstance
}

//	启动消息核心
//	初始化房间容器，加载已经持久化的房间等等
//
//	Author(Wind)
func (p *Service) Start() error {
	//	初始化房间容器对象
	p.container = GetRoomContainer()

	//	返回成功
	return nil
}

//	关闭消息核心
//	持久化房间容器等
//
//	Author(Wind)
func (p *Service) End() error {
	return nil
}
