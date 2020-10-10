package chat

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

//	Chat 首页控制器
//
//	Author(Wind)
func NewIndexController() *IndexController {
	indexOnceInstance.Do(func() {
		indexInstance = &IndexController{}
	})

	return indexInstance
}

//	首页模板控制器
//
//	Author(Wind)
func (p *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
