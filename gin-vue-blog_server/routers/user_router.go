package routers

import (
	"gin-vue-blog_server/api"
	"gin-vue-blog_server/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/email_login", userApi.EmailLoginView)                            // POST 上传（添加）
	router.GET("/user", middleware.JwtAuth(), userApi.UserListView)                // GET 查询
	router.PUT("/user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)    // 修改用户权限
	router.PUT("/user_password", middleware.JwtAuth(), userApi.UserUpdatePassword) // 修改用户密码
	router.DELETE("/user", middleware.JwtAdmin(), userApi.UserRemoveView)          // 删除用户 用的少
}
