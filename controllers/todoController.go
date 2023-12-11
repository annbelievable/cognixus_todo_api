package controllers

import (
	"cognixus_todo_api/inits"
	"cognixus_todo_api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(ctx *gin.Context) {
	var body struct {
		Title string
	}

	ctx.BindJSON(&body)

	if len(body.Title) == 0 {
		ctx.JSON(500, gin.H{"error": "Title is required."})
		return
	}

	todo := models.Todo{Title: body.Title, Complete: false}
	result := inits.DB.Create(&todo)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": todo})
}

func GetTodos(ctx *gin.Context) {
	var todos []models.Todo

	result := inits.DB.Find(&todos)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": todos})
}

func GetTodo(ctx *gin.Context) {

	var todo models.Todo

	result := inits.DB.First(&todo, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": todo})

}

func CompleteTodo(ctx *gin.Context) {
	var todo models.Todo

	result := inits.DB.First(&todo, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&todo).Updates(models.Todo{Complete: true})

	ctx.JSON(200, gin.H{"data": todo})
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	inits.DB.Delete(&models.Todo{}, id)

	ctx.JSON(200, gin.H{"data": "Todo has been deleted successfully"})
}
