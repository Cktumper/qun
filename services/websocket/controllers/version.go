package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type VersionController struct{}

var (
	versionInstance     *VersionController
	versionOnceInstance sync.Once
)

//	创建一个版本号控制器
//
//	Author(Wind)
func NewVersionController() *VersionController {
	versionOnceInstance.Do(func() {
		versionInstance = &VersionController{}
	})

	return versionInstance
}

//	版本号路由
//
//	Author(Wind)
func (p *VersionController) Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": gin.H{"version": "v0.0.1"}})
}
