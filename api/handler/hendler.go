// ginのhandlefuncを返す

package handler

import (
	"todo/model"
	"todo/usecase"

	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(usecase usecase.TodoUsecase) TodoHandler {
	return TodoHandler{usecase}
}

func (handler *TodoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := handler.usecase.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, tasks)
	}
}

func (handler *TodoHandler) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := strconv.Atoi(c.Param("id"))
		if ok != nil {
			c.JSON(500, gin.H{"error": ok.Error()})
			return
		}

		task, err := handler.usecase.GetOne(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, task)
	}
}

func (handler *TodoHandler) CreateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		task, err := handler.usecase.CreateOne(task)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, task)
	}
}

func (handler *TodoHandler) UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		task, err := handler.usecase.UpdateOne(task)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, task)
	}
}
