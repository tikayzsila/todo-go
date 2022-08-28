package main

import (
	"log"
	restgo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConf(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepositpry()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restgo.Server)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConf() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
