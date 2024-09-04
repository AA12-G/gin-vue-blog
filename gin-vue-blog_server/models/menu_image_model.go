package models

// MenuBannerModel 菜单图片关联表
type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
