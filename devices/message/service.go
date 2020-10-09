package message

import (
	"fmt"
	"sync"
)

type Service struct {
	buff chan Message
	size uint
}

//	默认消息缓冲通道
const CDefaultBufferSize = 1024

var (
	instance *Service
	once     sync.Once
)

//	构建消息服务
//
///	Author(Wind)
func NewService(size ...uint) *Service {
	once.Do(func() {
		instance = &Service{size: CDefaultBufferSize}

		if len(size) > 0 && size[0] > 0 {
			instance.size = size[0]
		}
	})

	return instance
}

//	获取服务实例
//
//	Author(Wind)
func GetService() *Service {
	return instance
}

//	发送信息
//
//	Author(Wind)
func (p *Service) Send(msg *Message) {
	//	执行发送
	go func() {
		p.buff <- *msg
	}()
}

//	接收消息
//
//	Author(Wind)
func (p *Service) Receiver(conn Connection) {
	go func() {
		for {
			//	接收消息
			_, message, err := conn.GetConnection().ReadMessage()
			if err != nil {
				continue
			}

			//	接收消息，处理逻辑
			fmt.Println(message)
		}
	}()
}

//	执行发送任务
//
//	Author(Wind)
func (p *Service) Run() {
	for {
		<-p.buff
	}
}

//	执行服务
//
//	Author(Wind)
func (p *Service) Start() error {
	p.buff = make(chan Message, p.size)

	//	启动服务发送任务
	go p.Run()

	return nil
}

//	关闭发送服务
//
//	Author(Wind)
func (p *Service) End() error {
	return nil
}
