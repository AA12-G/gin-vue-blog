package api

import "gin-vue-blog_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

// 实例化
var ApiGroupApp = new(ApiGroup)
