package routers

import (
	"gin-vue-blog_server/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouterGroup := router.Group("api")
	routerGroup := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	return router
}
