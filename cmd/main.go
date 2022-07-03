package main

import (
	"log"
	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error Server %s", err.Error())
	}
}
