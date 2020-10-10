package notify

import (
	"fmt"
	"log"
	"peon.top/qun/controllers/ws"
	"peon.top/qun/devices/session"
	"sync"
)

type Service struct {
	buff  chan Message
	size  uint
	route *Route
}

//	默认消息缓冲通道
const CDefaultBufferSize = 1

var (
	instance *Service
	once     sync.Once
)

//	构建消息服务
//
///	Author(Wind)
func NewService(size ...uint) *Service {
	once.Do(func() {
		instance = &Service{
			size:  CDefaultBufferSize,
			route: NewRoute(),
		}

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
func (p *Service) Send(msg ws.IMsg) {
	//	执行发送
	go func() {
		p.buff <- msg
	}()
}

//	接收消息
//
//	Author(Wind)
func (p *Service) Receiver(conn Connection) {
	go func() {
		for {
			//	接收消息
			msgType, message, err := conn.GetConnection().ReadMessage()
			if msgType < 0 || err != nil {
				break
			}

			//	接收消息，处理逻辑
			log.Println("接收数据：" + string(message))
			_ = p.route.Do(p, message)
		}
	}()
}

//	执行发送任务
//
//	Author(Wind)
func (p *Service) Run() {
	for {
		select {
		case message := <-p.buff:
			ss := message.GetTo().([]session.Session)

			for _, sess := range ss {
				fmt.Println(sess)
				func(sess Session) {
					if sess.GetConnection() == nil {
						return
					}
					_ = sess.GetConnection().WriteJSON(message.GetPacket())
				}(&sess)
			}
		default:

		}

	}
}

func (p *Service) loadRoutes() {
	p.route.Add(2, ws.NewMessagePacketController().Chat)
}

//	执行服务
//
//	Author(Wind)
func (p *Service) Start() error {
	p.buff = make(chan Message, p.size)

	//	加载消息路由
	p.loadRoutes()

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
