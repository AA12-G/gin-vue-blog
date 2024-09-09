package api

import (
	"gin-vue-blog_server/api/advert_api"
	"gin-vue-blog_server/api/images_api"
	"gin-vue-blog_server/api/menu_api"
	"gin-vue-blog_server/api/settings_api"
	"gin-vue-blog_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
}

// 实例化
var ApiGroupApp = new(ApiGroup)
