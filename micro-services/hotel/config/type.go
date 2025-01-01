package config

type Config struct {
	DB     DBConfig        `json:"db"`
	Server ServerConfig    `json:"server"`
	Redis  RedisConfig     `json:"redis"`
	Bank   BankGRPCConfig  `json:"bank"`
	Notif  NotifGRPCConfig `json:"notif"`
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
	HttpPort          uint   `json:"httpPort"`
	Secret            string `json:"secret"`
	AuthExpMinute     uint   `json:"authExpMin"`
	AuthRefreshMinute uint   `json:"authExpRefreshMin"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type BankGRPCConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type NotifGRPCConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
