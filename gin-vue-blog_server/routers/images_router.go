package routers

import "gin-vue-blog_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", imagesApi.ImagesUploadView)
}
