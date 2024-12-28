package config

type Config struct {
	DB               DBConfig         `json:"db"`
	Server           ServerConfig     `json:"server"`
	Redis            RedisConfig      `json:"redis"`
	HotelService     HotelService     `json:"hotelService"`
	TransportService TransportService `json:"transportService"`
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
	GrpcPort          uint   `json:"grpcPort"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type HotelService struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type TransportService struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
