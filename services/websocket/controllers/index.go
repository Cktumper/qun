package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"peon.top/qun/config"
	"strconv"
	"sync"
)

type IndexController struct{}

var (
	indexInstance     *IndexController
	indexOnceInstance sync.Once
)

//	创建一个版本号控制器
//
//	Author(Wind)
func NewIndexController() *IndexController {
	indexOnceInstance.Do(func() {
		indexInstance = &IndexController{}
	})

	return indexInstance
}

//	群聊器首页路由
//
//	Author(Wind)
func (p *IndexController) HTML(c *gin.Context) {
	//	获取用户ID信息
	userId, _ := strconv.Atoi(c.Query("user_id"))

	//	渲染模板
	c.HTML(http.StatusOK, "wss.html", gin.H{
		"Nickname": c.Query("nickname"),
		"UserId":   userId,
		"RoomId":   1,
		"Host": fmt.Sprintf("%s%s/wss/connect",
			config.ReadString("APP_WSS"),
			config.ReadString("APP_URL"),
		),
	})
}
