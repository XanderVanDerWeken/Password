package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xandervanderweken/Password/internal/controllers"
	"github.com/xandervanderweken/Password/internal/dependencies"
)

func main() {
	router := gin.Default()

	configureRouter(router, dependencies.PasswordController)

	router.Run()
}

func configureRouter(router *gin.Engine, passwordController controllers.PasswordController) {
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("", apiInfo)

		passwordGroup := apiGroup.Group("/passwords")
		{
			passwordGenerateGroup := passwordGroup.Group("/generate")
			{
				passwordGenerateGroup.POST("", passwordController.PostCreatePassword)
				passwordGenerateGroup.GET("", passwordController.GetCreatePassword)
			}
		}
	}
}

func apiInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"state":   "alive",
		"version": "v1.0",
	})
}
