package main

import (
	"log"
	restgo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepositpry()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restgo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
