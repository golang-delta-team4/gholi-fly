package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	HttpPort uint `json:"httpPort"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
