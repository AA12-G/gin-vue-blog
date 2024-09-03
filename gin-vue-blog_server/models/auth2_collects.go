package models

import "time"

// User2Collects 记录用户什么时候收藏了什么文章
type User2Collects struct {
	UserID    uint `gorm:"primaryKey"`
	UserModel uint `gorm:"foreignKey:UserID"`
	// 下面不会在表中生成对应字段的
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
