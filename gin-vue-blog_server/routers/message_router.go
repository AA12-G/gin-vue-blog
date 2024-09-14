package routers

import (
	"gin-vue-blog_server/api"
	"gin-vue-blog_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("/message", middleware.JwtAuth(), messageApi.MessageCreateView)       // POST 发送消息
	router.GET("/message_all", messageApi.MessageListAllView)                         // 查看所有消息列表(管理员)
	router.GET("/message", middleware.JwtAuth(), messageApi.MessageListView)          // 用户的消息
	router.GET("/message_record", middleware.JwtAuth(), messageApi.MessageRecordView) // 查看消息详情（聊天记录）
}
