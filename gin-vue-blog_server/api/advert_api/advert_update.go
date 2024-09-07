package advert_api

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 修改广告
// @Tags 广告管理
// @Summary 修改广告
// @Description 修改广告
// @Param data body  AdvertRequest  true "广告的一些参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
// 修改一个广告
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//广告存不存在的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	err = global.DB.Model(&advert).Updates(map[string]any{
		"title":   cr.Title,
		"href":    cr.Href,
		"images":  cr.Images,
		"is_show": cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.FailWithMessage("修改广告成功", c)
}
