package config

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id"`
	Key      string `yaml:"key" json:"key"`
	Redirect string `yaml:"redirect" json:"redirect"` // 登录之后的回调地址
}
