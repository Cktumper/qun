package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Service struct {
	//	公有方法
	EnableTLS bool

	//	私有方法
	engine *gin.Engine
	route  *Route
	port   int
}

//	构建一个WEB Socket服务
//
//	Author(Wind)
func NewService(port int) *Service {
	//	构建服务对象
	p := &Service{
		engine: gin.Default(),
		port:   port,
	}

	//	加载路由对象
	p.route = NewRoute(p.engine)

	//	返回自己
	return p
}

//	开启服务
//
//	Author(Wind)
func (p *Service) Start() error {
	//	加载路由表
	p.route.Load()

	//	判定是否需要使用 HTTPS
	if p.EnableTLS {
		return ErrNotSupportTLS
	}

	//	使用 HTTP 监听端口
	return p.engine.Run(fmt.Sprintf(":%d", p.port))
}

//	关闭服务
//
//	Author(Wind)
func (p *Service) End() error {
	return nil
}
