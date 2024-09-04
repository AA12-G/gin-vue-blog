package settings_api

import (
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// 系统管理视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "ok"})
	//res.Ok(map[string]string{}, "xxx", c)
	//res.OkWithData(map[string]string{
	//	"id": "1",
	//}, c)
	res.FailWithCode(1, c)
}
