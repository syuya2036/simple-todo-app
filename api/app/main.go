package main

import (
	"os"

	"todo/handler"
	"todo/repository"
	"todo/usecase"
)

func main() {
	db := repository.NewSqlHandler()
	repo := repository.NewRepository(db)
	usecase := usecase.NewTodoUsecase(repo)
	router := handler.Init()
	handler.TodoRouting(router, usecase)

	port := os.Getenv("PORT")
	hostName := os.Getenv("HOSTNAME")
	router.Run(hostName + ":" + port)
}
