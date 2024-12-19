package config

type Config struct {
	DB     DBConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	HttpHost          string `mapstructure:"httpHost"`
	HttpPort          uint   `mapstructure:"httpPort"`
	Secret            string `mapstructure:"secret"`
	AuthExpMinute     uint   `mapstructure:"authExpMin"`
	AuthRefreshMinute uint   `mapstructure:"authExpRefreshMin"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}
