package routers

import (
	"gin-vue-blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "123")
	})
	return router
}
