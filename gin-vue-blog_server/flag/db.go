package flag

import (
	"gin-vue-blog_server/global"
	"gin-vue-blog_server/models"
)

func MakeMigration() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "collectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	global.Log.Infof("开始迁移数据库")
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.BannerModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.FeedbackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[success] 生成数据库表结构成功")
}
