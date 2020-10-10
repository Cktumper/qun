package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"peon.top/qun/app/packet"
	"peon.top/qun/devices/message"
	"peon.top/qun/devices/session"
	"sync"
)

type LoginController struct {
	upgrader *websocket.Upgrader
}

var (
	loginInstance     *LoginController
	loginOnceInstance sync.Once
)

//	WSS 登陆控制器
//
//	Author(Wind)
func NewLoginController() *LoginController {
	loginOnceInstance.Do(func() {
		loginInstance = &LoginController{
			upgrader: &websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		}
	})

	return loginInstance
}

//	WSS Upgrade To Login
//
//	Author(Wind)
func (p *LoginController) Login(c *gin.Context) {
	//	判定连接是否是WSS升级请求
	//	若不是，则返回登陆失败
	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "错误的套接字"})
		c.Abort()
		return
	}

	//	否则启动 WSS
	conn, err := p.upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "升级权限失败"})
		c.Abort()
		return
	}

	//	创建 Session ,并绑定连接至 Session
	sess := session.NewSession()
	if err := sess.Create(conn); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "构建连接失败"})
		c.Abort()
		return
	}

	//	绑定 Session
	_ = session.GetService().Set(sess)

	//	监听消息接收事件
	message.GetService().Receiver(sess)

	//	发送成功消息包
	message.GetService().Send(message.NewMessage(
		&session.System,
		[]message.Session{sess},
		&packet.Connect{SessionId: sess.GetSessionId(), Flag: 1}),
	)
}
