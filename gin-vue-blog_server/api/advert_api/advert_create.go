package advert_api

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`   // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 图片
	IsShow *bool  `json:"is_show" binding:"required" msg:"请选择是否展示"`   // 是否展示
}

// AdvertCreateview 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
// 上传（添加）一个广告
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}

	res.FailWithMessage("添加广告成功", c)
}
