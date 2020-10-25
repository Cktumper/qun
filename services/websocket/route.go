package websocket

import (
	"github.com/gin-gonic/gin"
	"peon.top/qun/services/websocket/controllers"
)

type Route struct {
	engine *gin.Engine
}

//	构建一个路由对象
//
//	Author(Wind)
func NewRoute(engine *gin.Engine) *Route {
	return &Route{engine: engine}
}

func (p *Route) Load() {
	//	创建一个路由组
	r := p.engine.Group("wss")

	//	版本号控制器
	r.GET("version", controllers.NewVersionController().Version)
}
