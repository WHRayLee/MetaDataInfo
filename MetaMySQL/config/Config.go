package config

type ServerInfo struct {
	Version         string    `mapstructure:"Version"`
	Name            string    `mapstructure:"Name"`
	MySQLConfigInfo MySQLInfo `mapstructure:"MySQLInfo"`
}

type MySQLInfo struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Db       string `mapstructure:"DB"`
	Charset  string `mapstructure:"Charset"`
}
