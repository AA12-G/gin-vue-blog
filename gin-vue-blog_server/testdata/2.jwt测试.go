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
		UserID: 1,
		Role:   1,
		//Username: "zhangsan",
		NickName: "zs",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrX25hbWUiOiJ6cyIsInJvbGUiOjEsInVzZXJfaWQiOjEsImV4cCI6MTcyNTk0MjYxNC42MjgzMjQsImlzcyI6IjEyMzQifQ.KQrfxAgpRbbmsrh9QfDfStci3DcAMU0X-6IOG8O7aXY")
	fmt.Println(claims, err)
}
