package models

import "gin-vue-blog_server/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"` // 文章标题
	Abstract      string         `json:"abstract"`             // 文章摘要
	Content       string         `json:"content"`              // 文章内容
	LookCount     int            `json:"look_count"`           // 浏览量
	CommentCount  int            `json:"comment_count"`        // 评论量
	DiggCount     int            `json:"digg_count"`           // 点赞量
	CollectsCount int            `json:"collects_count"`
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  // 文章评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`                     // 文章作者
	UserID        uint           `json:"user_id"`                                        // 用户 	// 文章ID
	Category      string         `gorm:"size:20" json:"category"`                        // 文章分类
	Source        string         `json:"source"`                                         // 文章来源
	Link          string         `json:"link"`                                           // 文章来源链接
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`                   // 文章封面
	BannerID      uint           `json:"banner_id"`                                      // 文章封面ID
	NickName      string         `gorm:"size:42" json:"nick_name"`                       // 文章作者昵称
	BannerPath    string         `json:"banner_path"`                                    // 文章封面路径
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`                // 文章标签
}
