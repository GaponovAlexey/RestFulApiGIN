package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/handler"
	"github.com/gaponovalexey/todo-app/pkg/repository"
	"github.com/gaponovalexey/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//config LogrusJson
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//viper init
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initial config: %s", err.Error())
	}
	//env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	//db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		logrus.Fatal("failed to init", err)
	}

	//repository
	repository := repository.NewRepository(db)

	//service
	services := service.NewService(repository)

	//handler
	handler := handler.NewHandler(services)

	//start
	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error Server %s", err.Error())
		}
	}()

	logrus.Print("START todoApp")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("todoApp STOP")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error srv:%s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error db:%s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
