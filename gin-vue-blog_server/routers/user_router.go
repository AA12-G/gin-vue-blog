package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/user_email", userApi.EmailLoginView) // POST 上传（添加）
}
