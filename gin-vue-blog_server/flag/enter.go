package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string
}

// Parse 解析命令行： go run main.go -db
func Parse() Option {
	// go run main.go -db
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	// 解析命令
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsStopWeb 是否停止web项目
func IsStopWeb(option Option) bool {
	if option.DB {
		return true
	}
	return true // 停止web项目
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		// 迁移数据库
		MakeMigration()
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}
	sys_flag.Usage()
}
