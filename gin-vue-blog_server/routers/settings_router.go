package routers

import (
	"gin-vue-blog_server/api"
)

// 查看系统信息的路由
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingsApi.SettingsInfoView)
}
