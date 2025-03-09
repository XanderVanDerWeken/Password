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
	var newPasswordBody NewPassword
	context.BindJSON(&newPasswordBody)

	password, err := c.generatorService.CreateNewPassword(newPasswordBody.Length, newPasswordBody.UseNumbers, newPasswordBody.UseSpecials)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorDto{
			Message: "Could not create Password",
		})
	} else {
		context.JSON(http.StatusCreated, PasswordDto{
			Password: password,
		})
	}
}
