package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(configPath string) (Config, error) {
	viper.AddConfigPath(configPath)
    viper.SetConfigName("config")
    viper.SetConfigType("json")
    viper.AutomaticEnv()
	var config Config
    err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return  config, nil
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig(configPath)
	if err != nil {
		panic(err)
	}
	return c
}