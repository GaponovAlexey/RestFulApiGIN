package apiserver

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BindAddr string `env:"BindAddr" env-default:"3000"`
	LogLevel string `env:"logLevel" env-default:"debug"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			var helpText = "help"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Println(help)
		}
	})
	return instance
}
