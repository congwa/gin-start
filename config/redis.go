package config

type Redis struct {
	Addr       string `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 服务器地址:端口
	Password   string `mapstructure:"password" json:"password" yaml:"password"`       // 密码
	DB         int    `mapstructure:"db" json:"db" yaml:"db"`                         // 单实例模式下redis的哪个数据库
	UseCluster bool   `mapstructure:"useCluster" json:"useCluster" yaml:"useCluster"` // 是否使用集群模式
}
