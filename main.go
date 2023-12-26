package main

import (
	"cognixus_todo_api/controllers"
	"cognixus_todo_api/inits"
	"cognixus_todo_api/middlewares"
	"cognixus_todo_api/models"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
	inits.SetGithubOauthConfig()
	inits.SetGoogleOauthConfig()
}

func main() {
	inits.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	cookieSecret := os.Getenv("COOKIE_SECRET")
	store := sessions.NewCookieStore([]byte(cookieSecret))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/login-required/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login Required!",
		})
	})

	r.GET("/login/google/", controllers.HandleGoogleLogin)
	r.GET("/callback/google/", controllers.HandleGoogleCallback)

	r.GET("/login/github/", controllers.HandleGithubLogin)
	r.GET("/callback/github/", controllers.HandleGithubCallback)

	r.POST("/api/todo/", middlewares.AuthMiddleware(), controllers.CreateTodo)
	r.GET("/api/todo/", middlewares.AuthMiddleware(), controllers.GetTodos)
	r.GET("/api/todo/:id", middlewares.AuthMiddleware(), controllers.GetTodo)
	r.PUT("/api/todo/complete/:id", middlewares.AuthMiddleware(), controllers.CompleteTodo)
	r.DELETE("/api/todo/:id", middlewares.AuthMiddleware(), controllers.DeleteTodo)

	r.Run()
}
