package apiserver

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BindAddr string `toml:"bind_addr`
	LogLevel string `toml:"log_level`
}

// NewConfig...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		LogLevel: "debug",
	}
}

func init() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initial config: %s", err.Error()) // инициализация
	}
	var s Config
	s.BindAddr = viper.GetString("config.BindAddr")
	s.LogLevel = viper.GetString("config.LogLevel")
}

func initConfig() error {
	viper.AddConfigPath("configs") // путь
	viper.SetConfigName("config")  // файл
	return viper.ReadInConfig()
}
