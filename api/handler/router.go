// router.go

package handler

import (
	"todo/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	return router
}

func TodoRouting(router *gin.Engine, usecase usecase.TodoUsecase) {
	handler := NewTodoHandler(usecase)
	router.GET("/tasks", handler.GetAll())
	router.GET("/tasks/:id", handler.GetOne())
	router.POST("/tasks", handler.CreateOne())
	router.PUT("/tasks/:id", handler.UpdateOne())
}