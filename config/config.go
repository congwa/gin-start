package config

type Config struct {
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Server  Server  `mapstructure:"server" json:"server" yaml:"server"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
