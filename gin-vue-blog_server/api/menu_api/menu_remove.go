package menu_api

import (
	"fmt"
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.AgrumentError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}