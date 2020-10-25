package controllers

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type WSController struct{}

var (
	wsInstance     *WSController
	wsOnceInstance sync.Once
)

//	构建一个新的 Websocket 初始化连接控制器
//	这里也是用户进来等的初始位置
//
//	Author(Wind)
func NewWSController() *WSController {
	wsOnceInstance.Do(func() {
		wsInstance = &WSController{}
	})

	return wsInstance
}

//	将 HTTP 连接升级成 Websocket 连接
//	并创建房间用户，将用户加入至房间中
//	并监听用户端口收消息
//
//	Author(Wind)
func (p *WSController) Upgrade(c *gin.Context) {

}
