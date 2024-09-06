package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", imagesApi.ImagesUploadView)  // POST 上传（添加）
	router.GET("/images", imagesApi.ImageListView)      // GET 查询
	router.DELETE("/images", imagesApi.ImageRemoveView) // DELETE 删除
	router.PUT("/images", imagesApi.ImageUpdateView)    // PUT 修改
}
