package settings_api

import "github.com/gin-gonic/gin"

// 系统管理视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}
