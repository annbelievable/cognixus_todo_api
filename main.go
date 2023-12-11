package main

import (
	"cognixus_todo_api/controllers"
	"cognixus_todo_api/inits"
	"cognixus_todo_api/models"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/api/todo/", controllers.CreateTodo)
	r.GET("/api/todo/", controllers.GetTodos)
	r.GET("/api/todo/:id", controllers.GetTodo)
	r.PUT("/api/todo/complete/:id", controllers.CompleteTodo)
	r.DELETE("/api/todo/:id", controllers.DeleteTodo)

	r.Run()
}
