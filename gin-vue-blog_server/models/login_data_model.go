package models

// LoginDataModel 用户登录信息表
type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` // 登录的IP
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登陆设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
