package main

import (
	"log"

	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/handler"
	repository "github.com/gaponovalexey/todo-app/pkg/repo"
	"github.com/gaponovalexey/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error Server %s", err.Error())
	}
}
