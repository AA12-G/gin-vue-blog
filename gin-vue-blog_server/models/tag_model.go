package models

// TagModel 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`           // 标签名
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"_"` // 标签对应的文章
}
