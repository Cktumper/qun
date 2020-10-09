package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

//	控制器结构
type InformationController struct {
	name, version string
}

//	单例对象
var (
	informationInstance     *InformationController
	informationOnceInstance sync.Once
)

//	创建一个新的系统信息控制器
//
//	Author(Wind)
func NewInformationController() *InformationController {
	//	生成控制器单例
	informationOnceInstance.Do(func() {
		informationInstance = &InformationController{
			name:    "Qun",
			version: "0.0.1",
		}
	})

	//	初始化控制器
	return informationInstance
}

//	获取系统版本号
//
//	Author(Wind)
func (p *InformationController) Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": fmt.Sprintf("%s %s", p.name, p.version)})
}

//	Echo 控制器
//
//	Author(Wind)
func (p *InformationController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": "pong"})
}
