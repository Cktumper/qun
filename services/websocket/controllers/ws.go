package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"peon.top/qun/chat"
	"peon.top/qun/services/websocket/app"
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
	var request app.UpgradeRequest

	//	绑定请求信息
	if err := request.Bind(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "连接失败", "data": nil})
		c.Abort()
		return
	}

	//	判定连接是否为 Websocket 连接
	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "非法的连接", "data": nil})
		c.Abort()
		return
	}

	//	查找房间号
	room, err := chat.GetRoomContainer().Get(request.RoomId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error(), "data": nil})
		c.Abort()
		return
	}

	//	构建升级器，并允许跨域连接
	up := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//	创建一个连接
	wsConn, err := up.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "升级连接失败", "data": nil})
		c.Abort()
		return
	}

	//	构建一个 Session
	session := app.NewSession(
		app.NewUser(request.UserId, request.Nickname),
		app.NewConnection(wsConn),
	)

	//	将 Session 加入到 Chat 聊天室服务中
	if err := room.Join(session); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error(), "data": nil})
		c.Abort()
		return
	}
}
