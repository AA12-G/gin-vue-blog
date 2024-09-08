package menu_api

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
	"gin-vue-blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banner []Banner `json:"banners"`
}

func (MenuApi) MenuFindView(c *gin.Context) {
	// 先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort asc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		// model就是一个菜单
		// 解决null问题
		var banners = make([]Banner, 0)
		//banners:= []Banner{}
		//var banners []Banner
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banner:    banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
