package main

import (
	"fmt"
	"gin-vue-blog_server/core"
	"gin-vue-blog_server/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
}
