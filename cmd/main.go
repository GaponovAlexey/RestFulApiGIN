package main

import (
	"log"

	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/handler"
	"github.com/gaponovalexey/todo-app/pkg/repository"
	"github.com/gaponovalexey/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initial config: %s", err.Error())
	}
	//repository
	repository := repository.NewRepository()
	
	//service
	services := service.NewService(repository)
	
	//handler
	handler := handler.NewHandler(services)

	//start
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("3000"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error Server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
