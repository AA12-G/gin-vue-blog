package settings_api

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type SettingUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 查询某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.AgrumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qinqiu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}

}
