package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Upload   Upload   `yaml:"upload"`
	QQ       QQ       `canGet:"qq" yaml:"qq"`
	Email    Email    `canGet:"email" yaml:"email"`
	Jwt      Jwt      `canGet:"jwt" yaml:"jwt"`
	QiNiu    QiNiu    `canGet:"qiNiu" yaml:"qi_niu"`
}
