package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.POST("/tag", tagApi.TagCreateView)    // POST 上传（添加）
	router.GET("/tag", tagApi.TagFindView)       // GET 查询广告列表
	router.PUT("/tag/:id", tagApi.TagUpdateView) // PUT 修改广告
	router.DELETE("/tag", tagApi.TagRemoveView)  // DELETE 删除广告
}
