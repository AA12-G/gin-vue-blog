package jwts

import (
	"errors"
	"fmt"
	"gin-vue-blog_server/global"
	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid claim type")

}
