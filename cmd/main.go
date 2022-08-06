package main

import (
	// "database/sql"
	"os"

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
	logrus.SetFormatter(new(logrus.JSONFormatter))
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
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error Server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
