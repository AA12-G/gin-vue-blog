package routers

import (
	"gin-vue-blog_server/api"
	"gin-vue-blog_server/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/user_email", userApi.EmailLoginView)              // POST 上传（添加）
	router.GET("/user", middleware.JwtAuth(), userApi.UserListView) // GET 查询
}
