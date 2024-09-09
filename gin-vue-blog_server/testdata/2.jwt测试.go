package main

import (
	"fmt"
	"gin-vue-blog_server/core"
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "zhangsan",
		NickName: "zs",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InpoYW5nc2FuIiwibmlja19uYW1lIjoienMiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE3MjU4Nzk0MDMuNTk2NTcyOSwiaXNzIjoiMTIzNCJ9.wx6hlZyL0cVd1WxZ08NU_-PF_PO_KG5")
	fmt.Println(claims, err)
}
