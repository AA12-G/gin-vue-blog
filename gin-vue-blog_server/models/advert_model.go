package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` // 广告标题 唯一
	Href   string `gorm:"json:"href"`           // 跳转链接
	Images string `son:"images"`                // 图片
	IsShow *bool  `json:"is_show"`              // 是否显示
}
