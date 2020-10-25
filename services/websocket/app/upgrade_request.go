package app

import "github.com/gin-gonic/gin"

type UpgradeRequest struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	RoomId   int    `form:"room_id" json:"room_id" binding:"required"`
}

//	使用 Gin 绑定升级请求
//
//	Author(Wind)
func (p *UpgradeRequest) Bind(c *gin.Context) (err error) {
	if err = c.ShouldBind(p); err == nil {
		return nil
	}

	return err
}
