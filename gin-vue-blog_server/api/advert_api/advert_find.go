package advert_api

import (
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/res"
	"gin-vue-blog_server/service/common"
	"github.com/gin-gonic/gin"
)

// 查询广告列表
func (AdvertApi) AdvertFindView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.AgrumentError, c)
		return
	}
	list, count, _ := common.ComList(models.AdvertModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
