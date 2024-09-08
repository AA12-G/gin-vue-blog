package models

import "gin-vue-blog_server/models/ctype"

type ImageSort struct {
	ImageID uint `json:"image_id"` // 已经上传了的图片，对应file
	Sort    int  `json:"sort"`
}

type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`
	Path         string        `gorm:"size:32" json:"path"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` // 菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                               // 菜单的切换时间 0表示不切换
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       // 菜单顺序
}
