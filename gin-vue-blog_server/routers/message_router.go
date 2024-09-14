package routers

import (
	"gin-vue-blog_server/api"
	"gin-vue-blog_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("/message", middleware.JwtAuth(), messageApi.MessageCreateView) // POST 上传（添加）
}
