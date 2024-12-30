package config

type Config struct {
	DB       DBConfig       `json:"db"`
	Server   ServerConfig   `json:"server"`
	SMTP     SMTPConfig     `json:"smtp"`
	NATS     NATSConfig     `json:"nats"`
	UserGRPC UserGRPCConfig `json:"userGrpc"`
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
	HttpPort uint   `json:"httpPort"`
	GRPCPort uint   `json:"grpcPort"`
	Secret   string `mapstructure:"secret"`
}

type SMTPConfig struct {
	Sender   string `json:"sender"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NATSConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type UserGRPCConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
