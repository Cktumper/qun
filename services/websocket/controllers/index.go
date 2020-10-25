package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.HTML(http.StatusOK, "wss.html", gin.H{"Nickname": c.Query("nickname")})
}
