package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("/menu", menuApi.MenuCreateView)    // POST 上传（添加）
	router.GET("/menu", menuApi.MenuFindView)       // GET 查询菜单
	router.GET("/menu_name", menuApi.MenuNameList)  // GET 查询菜单
	router.PUT("/menu/:id", menuApi.MenuUpdateView) // PUT 修改菜单
	router.DELETE("/menu", menuApi.MenuRemoveView)  // DELETE 删除菜单
	router.GET("/menu/:id", menuApi.MenuDetailView) // 查看菜单详情
}
