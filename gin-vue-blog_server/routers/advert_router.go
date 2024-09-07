package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	router.POST("/advert", advertApi.AdvertCreateView)    // POST 上传（添加）
	router.GET("/advert", advertApi.AdvertFindView)       // GET 查询广告列表
	router.PUT("/advert/:id", advertApi.AdvertUpdateView) // PUT 修改广告
	router.DELETE("/advert", advertApi.AdvertRemoveView)  // DELETE 删除广告
}
