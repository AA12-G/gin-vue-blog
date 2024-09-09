package jwts

import (
	"gin-vue-blog_server/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GenToken 生成token
func GenToken(user JwtPayLoad) (string, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //默认两小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)

}
