package http

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"peon.top/qun/config"
	"peon.top/qun/controllers/chat"
	"peon.top/qun/controllers/system"
)

type Http struct {
	engine *gin.Engine
}

func NewHTTP() *Http {
	return &Http{
		engine: gin.Default(),
	}
}

//	获取服务实例
func (p *Http) GetService() *gin.Engine {
	return p.engine
}

//	加载 HTTP 路由
func (p *Http) LoadRouter() *Http {
	//	加载系统路由
	p.loadSystemRouter()

	p.engine.GET("/chat/login", chat.NewLoginController().Login)

	return p
}

//	加载系统路由
func (p *Http) loadSystemRouter() {
	router := p.engine.Group("info")

	router.GET("version", system.NewInformationController().Version)
	router.GET("ping", system.NewInformationController().Ping)
}

//	监听 HTTP 服务
func (p *Http) Listen() error {
	return endless.ListenAndServe(fmt.Sprintf("%s:%s", config.ReadString("SERVICE_HOST"), config.ReadString("SERVICE_PORT")), p.engine)
}

//	Start
func (p *Http) Start() error {
	p.LoadRouter()
	return p.Listen()
}

//	Close
func (p *Http) End() error {
	return nil
}
