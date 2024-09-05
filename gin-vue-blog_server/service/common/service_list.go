package common

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//默认没有日志
	DB := global.DB
	if option.Debug {
		//如果有bug才打印
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	//图片的列表排序
	if option.Sort == "" {
		//option.Sort = "id desc" //默认按照id倒序
		option.Sort = "id asc" //默认按照id正序
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
