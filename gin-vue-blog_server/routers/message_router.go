package routers

import (
	"gin-vue-blog_server/api"
	"gin-vue-blog_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("/message", middleware.JwtAuth(), messageApi.MessageCreateView) // POST 发送消息
	router.GET("/message_all", messageApi.MessageListAllView)                   // 查看所有消息列表
	router.GET("/message", middleware.JwtAuth(), messageApi.MessageListView)
}
