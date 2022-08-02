package main

import (
	"log"

	"github.com/gaponovalexey/go-restapi/pkg/app/apiserver"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initial config: %s", err.Error()) // инициализация
	}
	
	port, level, db := viper.GetString("port"), viper.GetString("level"), viper.GetString("db")
	config := apiserver.NewConfig(port, level, db)
	
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
func initConfig() error {
	viper.AddConfigPath("configs") // путь
	viper.SetConfigName("config")  // файл
	return viper.ReadInConfig()
}
