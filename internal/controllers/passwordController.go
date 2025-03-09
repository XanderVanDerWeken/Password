package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xandervanderweken/Password/internal/services"
)

type passwordController struct {
	generatorService services.GeneratorService
}

func NewPasswordController(
	generatorService services.GeneratorService) PasswordController {
	return &passwordController{
		generatorService: generatorService,
	}
}

func (c *passwordController) CreatePassword(context *gin.Context) {
	password := c.generatorService.CreateNewPassword()

	context.JSON(http.StatusCreated, PasswordDto{
		Password: password,
	})
}
