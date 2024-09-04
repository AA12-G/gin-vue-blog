package main

import (
	"fmt"
	"gin-vue-blog_server/core"
	"gin-vue-blog_server/flag"
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	option := flag.Parse()
	fmt.Println(option)
	if flag.IsStopWeb(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server 运行在 %s", addr)
	router.Run(addr) //端口默认8080
}
