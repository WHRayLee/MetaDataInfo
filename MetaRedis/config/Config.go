package config

type ServerInfo struct {
	Name            string    `mapstructure:"Name"`
	RedisConfigInfo RedisInfo `mapstructure:"RedisInfo"`
}

type RedisInfo struct {
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	DB           int    `mapstructure:"DB"`
	DialTimeout  int64  `mapstructure:"DialTimeout"`
	ReadTimeout  int64  `mapstructure:"ReadTimeout"`
	WriteTimeout int64  `mapstructure:"WriteTimeout"`
}
