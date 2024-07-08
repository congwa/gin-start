package config

type Config struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
}
