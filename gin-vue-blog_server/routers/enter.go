package routers

import (
	"gin-vue-blog_server/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")
	routerGroup := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	routerGroup.AdvertRouter()
	routerGroup.MenuRouter()
	routerGroup.UserRouter()
	return router
}
