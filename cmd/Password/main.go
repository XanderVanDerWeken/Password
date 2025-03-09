package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xandervanderweken/Password/internal/controllers"
	"github.com/xandervanderweken/Password/internal/services"
)

func main() {
	// Get Instances
	generatorService := services.NewGeneatorService()
	passwordController := controllers.NewPasswordController(generatorService)

	router := gin.Default()

	configureRouter(router, passwordController)

	router.Run()
}

func configureRouter(router *gin.Engine, passwordController controllers.PasswordController) {
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiGroup := router.Group("/api")
	{
		passwordGroup := apiGroup.Group("/passwords")
		{
			passwordGroup.POST("", passwordController.CreatePassword)
		}
	}
}
