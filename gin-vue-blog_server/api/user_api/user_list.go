package user_api

import (
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/ctype"
	"gin-vue-blog_server/models/res"
	"gin-vue-blog_server/service/common"
	"gin-vue-blog_server/utils/desens"
	"gin-vue-blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.AgrumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Email = desens.DesensitizationEmail(user.Email)
		user.Tel = desens.DesensitizationTel(user.Tel)
		// 脱敏
		users = append(users, user)
	}
	res.OkWithList(list, count, c)
}
